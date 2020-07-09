/*
Copyright © 2020 Adrian Pareja <adriancc5.5@gmail.com>
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"math/big"
	"net/http"
	"net/http/httputil"
	"net/url"
	"io/ioutil"
	"bytes"
	"github.com/spf13/viper"
	"github.com/lacchain/gas-relay-signer/rpc"
	"github.com/lacchain/gas-relay-signer/util"
	"github.com/lacchain/gas-relay-signer/model"
	"github.com/lacchain/gas-relay-signer/service"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var config *model.Config
var relaySignerService *service.RelaySignerService

func main() {
	config = getConfigFromFile()

	relaySignerService = new(service.RelaySignerService)
	relaySignerService.Init(config)
	setupRoutes()
}

func signTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Body:", r.Body)

	buf, _ := ioutil.ReadAll(r.Body)
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))

	fmt.Println("Request body : ", rdr1)
	fmt.Println("Request body : ", rdr2)

	var rpcMessage rpc.JsonrpcMessage

	err := json.NewDecoder(rdr1).Decode(&rpcMessage)
	if err != nil {
		fmt.Println("Error Decoding Json RPC")
		fmt.Println(err)
		return
	}

	fmt.Println("JSON-RPC Method:",rpcMessage.Method);

	if (rpcMessage.IsRawTransaction()){
		fmt.Println("Is a rawTransaction")
		var params []string
		err = json.Unmarshal(rpcMessage.Params, &params)
		if err != nil {
			fmt.Println("Error Unmarshaling Json RPC Params")
			fmt.Println(err)
		}

		//fmt.Println("Param 0:",params[0])

		decodeTransaction,_ := util.GetTransaction(params[0][2:])

		message, err := decodeTransaction.AsMessage(types.NewEIP155Signer(decodeTransaction.ChainId()))
    	if err != nil {
        	log.Fatal(err)
    	}

		//fmt.Println("Decode Transaction:",*decodeTransaction)
		fmt.Println("From:",message.From().Hex())
		if (decodeTransaction.To() != nil){
			fmt.Println("To:",decodeTransaction.To().Hex())
		}
		fmt.Println("Data:",hexutil.Encode(decodeTransaction.Data()))
		fmt.Println("GasLimit:",decodeTransaction.Gas())
		fmt.Println("Nonce",decodeTransaction.Nonce())
		fmt.Println("GasPrice:",decodeTransaction.GasPrice())
		fmt.Println("Value:",decodeTransaction.Value())
		v,r,s := decodeTransaction.RawSignatureValues();

		fmt.Println(fmt.Sprintf("Signature R %x",r))
		fmt.Println(fmt.Sprintf("Signature S %x",s))
		fmt.Println(fmt.Sprintf("Signature V %x",v))

		signature,_ := util.SignPayload(relaySignerService.Config.Application.Key, message.From().Hex(), decodeTransaction.To(), decodeTransaction.Data(), decodeTransaction.Gas(), decodeTransaction.Nonce())
		fmt.Println("signature:",signature)

		//relaySignerService := new(service.RelaySignerService)

		response := relaySignerService.SendMetatransaction(rpcMessage.ID, message.From(), decodeTransaction.To(), decodeTransaction.Data(), new(big.Int).SetUint64(decodeTransaction.Gas()), new(big.Int).SetUint64(decodeTransaction.Nonce()), signature)
		data, err := json.Marshal(response)
		w.Write(data)
	}else{
		r.Body=rdr2
		fmt.Println("Is not a rawTransaction")
		
		serveReverseProxy(config.Application.NodeURL,w,r)
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

func getConfigFromFile()(*model.Config){
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}
	var c model.Config
	if err := v.Unmarshal(&c); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}
	fmt.Printf("smartContract=%s AgentKey=%s\n", c.Application.ContractAddress, c.KeyStore.Agent)
	return &c
}

func setupRoutes() {
	fmt.Println("Init RelaySigner")
	http.HandleFunc("/", signTransaction)
	http.ListenAndServe(":9001", nil)
}