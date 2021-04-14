package controller

import (
	"errors"
	"math/big"
	"encoding/json"
	"encoding/hex"
	"fmt"
	"sync"
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
	"github.com/ethereum/go-ethereum/rlp"
	
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

		if decodeTransaction.Gas() > controller.RelaySignerService.GetGasLimit() {
			err := errors.New("transaction gas limit exceeds block gas limit") 
			data := handleError(rpcMessage.ID, err)
			w.Write(data)
			return
		}

		log.GeneralLogger.Println("From:",message.From().Hex())
		if (decodeTransaction.To() != nil){
			log.GeneralLogger.Println("To:",decodeTransaction.To().Hex())
		}
		log.GeneralLogger.Println("Data:",hexutil.Encode(decodeTransaction.Data()))
		log.GeneralLogger.Println("GasLimit:",decodeTransaction.Gas())
		log.GeneralLogger.Println("Nonce",decodeTransaction.Nonce())
		log.GeneralLogger.Println("GasPrice:",decodeTransaction.GasPrice())
		log.GeneralLogger.Println("Value:",decodeTransaction.Value())
		v,rInt,sInt := decodeTransaction.RawSignatureValues();

		log.GeneralLogger.Println(fmt.Sprintf("Signature R %064x",rInt))
		log.GeneralLogger.Println(fmt.Sprintf("Signature S %064x",sInt))
		log.GeneralLogger.Println(fmt.Sprintf("Signature V %x",v))

		var r [32]byte
		var s [32]byte
		rBytes,_ :=hex.DecodeString(fmt.Sprintf("%064x",rInt))
		sBytes,_ :=hex.DecodeString(fmt.Sprintf("%064x",sInt))

		copy(r[:],rBytes)
		fmt.Println(rBytes)
		fmt.Println(r)
		copy(s[:],sBytes)
		

		log.GeneralLogger.Println("senderSignature:",fmt.Sprintf("%064x",rInt)+fmt.Sprintf("%064x",sInt)+fmt.Sprintf("%x",v))

		signature, err := util.SignPayload(controller.RelaySignerService.Config.Application.Key, message.From().Hex(), decodeTransaction.To(), decodeTransaction.Data(), decodeTransaction.Gas(), decodeTransaction.Nonce())
		if err != nil {
			data := handleError(rpcMessage.ID, err)
			w.Write(data)
			return
		}

		log.GeneralLogger.Println("signature:",signature)

		var signingDataTx *model.RawTransaction

		if (decodeTransaction.To() != nil){
			signingDataTx = model.NewTransaction(decodeTransaction.Nonce(), *decodeTransaction.To(), decodeTransaction.Value(), decodeTransaction.Gas(), decodeTransaction.GasPrice(), decodeTransaction.Data())
		}else{
			signingDataTx = model.NewContractCreation(decodeTransaction.Nonce(), decodeTransaction.Value(), decodeTransaction.Gas(), decodeTransaction.GasPrice(), decodeTransaction.Data())
		}

		signingDataRLP, err := rlp.EncodeToBytes(signingDataTx.Data)
		if err != nil {
			fmt.Println("encode error: ", err)
		}

		log.GeneralLogger.Println("SigningDataRLP:",hexutil.Encode(signingDataRLP))

		lock.Lock()
		response := controller.RelaySignerService.SendMetatransaction(rpcMessage.ID, decodeTransaction.To(), signingDataRLP, uint8(v.Uint64()), r, s)
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
	}else if(rpcMessage.IsGetBlockByNumber()){
		log.GeneralLogger.Println("Is getBlockByNumber")
		var params []interface{}
		err = json.Unmarshal(rpcMessage.Params, &params)
		if err != nil {
			data := handleError(rpcMessage.ID, err)
			w.Write(data)
			return
		}

		var blockNumber *big.Int

		if params[0].(string)[0:2] == "0x"{
			number, err := hexutil.DecodeUint64(params[0].(string))
			if err != nil {
				data := handleError(rpcMessage.ID, err)
				w.Write(data)
				return
			}
	
			blockNumber = new(big.Int).SetUint64(number)
		}else if (params[0].(string) == "earliest"){
			blockNumber = new(big.Int).SetUint64(0)
		}

		response := controller.RelaySignerService.GetBlockByNumber(rpcMessage.ID,blockNumber)
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