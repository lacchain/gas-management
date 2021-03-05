package controller

import (
	"encoding/json"
	"encoding/hex"
	"fmt"
	"sync"
	"math/big"
	"net/http/httputil"
	"net/http"
	"net/url"
	"io/ioutil"
	"bytes"
	"github.com/lacchain/gas-relay-signer/rpc"
	"github.com/lacchain/gas-relay-signer/util"
	"github.com/lacchain/gas-relay-signer/model"
	"github.com/lacchain/gas-relay-signer/service"
	log "github.com/lacchain/gas-relay-signer/util"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var lock sync.Mutex

//RelayController is the main controller
type RelayController struct {
	// The controller's configuration
	Config *model.Config
	RelaySignerService *service.RelaySignerService
}

//Init controller
func (controller *RelayController) Init(config *model.Config, relaySignerService *service.RelaySignerService){
	controller.Config = config
	controller.RelaySignerService = relaySignerService
}

//SignTransaction ...
func (controller *RelayController) SignTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.GeneralLogger.Println("Body:", r.Body)

	buf, _ := ioutil.ReadAll(r.Body)
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))

	log.GeneralLogger.Println("Request body : ", rdr1)

	var rpcMessage rpc.JsonrpcMessage

	err := json.NewDecoder(rdr1).Decode(&rpcMessage)
	if err != nil {
		log.GeneralLogger.Println("Error Decoding Json RPC")
		log.GeneralLogger.Println(err)
		return
	}

	log.GeneralLogger.Println("JSON-RPC Method:",rpcMessage.Method);

	if (rpcMessage.IsPrivTransaction()){
		r.Body=rdr2
		log.GeneralLogger.Println("Is a private Transaction, forward to Besu->Orion")

		serveReverseProxy(controller.Config.Application.NodeURL,w,r)
	}else if (rpcMessage.IsPrivRawTransaction()){
		r.Body=rdr2
		log.GeneralLogger.Println("Is a private send Transaction, decrease gas used")

		controller.RelaySignerService.DecreaseGasUsed()

		log.GeneralLogger.Println("forward to Besu->Orion")
		serveReverseProxy(controller.Config.Application.NodeURL,w,r)
	}else if (rpcMessage.IsRawTransaction()){
		log.GeneralLogger.Println("Is a rawTransaction")
		var params []string
		err = json.Unmarshal(rpcMessage.Params, &params)
		if err != nil {
			data := handleError(rpcMessage.ID, err)
			w.Write(data)
			return
		}

		decodeTransaction, err := util.GetTransaction(params[0][2:])
		if err != nil {
        	data := handleError(rpcMessage.ID, err)
			w.Write(data)
			return
    	}

		message, err := decodeTransaction.AsMessage(types.NewEIP155Signer(decodeTransaction.ChainId()))
    	if err != nil {
        	data := handleError(rpcMessage.ID, err)
			w.Write(data)
			return
    	}

		log.GeneralLogger.Println("From:",message.From().Hex())
		if (decodeTransaction.To() != nil){
			log.GeneralLogger.Println("To:",decodeTransaction.To().Hex())

		/*	if (relaySignerService.IsTargetPermitted(decodeTransaction.To().Hex())){
				relaySignerService.DecreaseGasUsed()
				log.GeneralLogger.Println("Old Smart Contract --> forward to Besu")
				serveReverseProxy(config.Application.NodeURL,w,r)
			} */
		}
		log.GeneralLogger.Println("Data:",hexutil.Encode(decodeTransaction.Data()))
		log.GeneralLogger.Println("GasLimit:",decodeTransaction.Gas())
		log.GeneralLogger.Println("Nonce",decodeTransaction.Nonce())
		log.GeneralLogger.Println("GasPrice:",decodeTransaction.GasPrice())
		log.GeneralLogger.Println("Value:",decodeTransaction.Value())
		v,r,s := decodeTransaction.RawSignatureValues();

		log.GeneralLogger.Println(fmt.Sprintf("Signature R %064x",r))
		log.GeneralLogger.Println(fmt.Sprintf("Signature S %064x",s))
		log.GeneralLogger.Println(fmt.Sprintf("Signature V %x",v))

		log.GeneralLogger.Println("senderSignature:",fmt.Sprintf("%064x",r)+fmt.Sprintf("%064x",s)+fmt.Sprintf("%x",v))

		senderSignature, err := hex.DecodeString(fmt.Sprintf("%064x",r)+fmt.Sprintf("%064x",s)+fmt.Sprintf("%x",v))

		if err != nil {
			data := handleError(rpcMessage.ID, err)
			w.Write(data)
			return
		}

		signature, err := util.SignPayload(controller.RelaySignerService.Config.Application.Key, message.From().Hex(), decodeTransaction.To(), decodeTransaction.Data(), decodeTransaction.Gas(), decodeTransaction.Nonce())
		if err != nil {
			data := handleError(rpcMessage.ID, err)
			w.Write(data)
			return
		}

		log.GeneralLogger.Println("signature:",signature)

		lock.Lock()
		response := controller.RelaySignerService.SendMetatransaction(rpcMessage.ID, message.From(), decodeTransaction.To(), decodeTransaction.Data(), new(big.Int).SetUint64(decodeTransaction.Gas()), new(big.Int).SetUint64(decodeTransaction.Nonce()), signature,senderSignature)
		lock.Unlock()
		data, err := json.Marshal(response)
		w.Write(data)
	}else if (rpcMessage.IsGetTransactionReceipt()){
		r.Body=rdr2
		log.GeneralLogger.Println("Is getTransactionReceipt")
		var params []string
		err = json.Unmarshal(rpcMessage.Params, &params)
		if err != nil {
			data := handleError(rpcMessage.ID, err)
			w.Write(data)
			return
		}
		response := controller.RelaySignerService.GetTransactionReceipt(rpcMessage.ID,params[0][2:])
		data, _ := json.Marshal(response)
		w.Write(data)
		return

	}else if(rpcMessage.IsGetTransactionCount()){
		log.GeneralLogger.Println("Is getTransactionCount")
		var params []string
		err = json.Unmarshal(rpcMessage.Params, &params)
		if err != nil {
			data := handleError(rpcMessage.ID, err)
			w.Write(data)
			return
		}
		response := controller.RelaySignerService.GetTransactionCount(rpcMessage.ID,params[0])
		data, _ := json.Marshal(response)
		w.Write(data)
		return
	}else{
		r.Body=rdr2
		log.GeneralLogger.Println("Is another type of transaction, reverse proxy")
		
		serveReverseProxy(controller.Config.Application.NodeURL,w,r)
	}
}

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

func handleError(messageID json.RawMessage, err error) ([]byte) {
	log.GeneralLogger.Println(err)
	data, _ := json.Marshal(service.HandleErrorRPCMessage(messageID,err))
	
	return data
}