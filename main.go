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
	"math/big"
	//"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	//"time"
//	"sync"
"io/ioutil"
"bytes"
	"github.com/lacchain/gas-relay-signer/rpc"
	"github.com/lacchain/gas-relay-signer/util"
	"github.com/lacchain/gas-relay-signer/service"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
	
	setupRoutes()

	fmt.Println("Terminating Program")
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

	fmt.Println("value:",rpcMessage.Method);

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

		//fmt.Println("Decode Transaction:",*decodeTransaction)
		fmt.Println("From:0x63949701cd0e1cc04dfea0afbf410968f10ff4b6")
		fmt.Println("To:",decodeTransaction.To().Hex())
		fmt.Println("Data:",hexutil.Encode(decodeTransaction.Data()))
		fmt.Println("GasLimit:",decodeTransaction.Gas())
		fmt.Println("Nonce",decodeTransaction.Nonce())
		fmt.Println("GasPrice:",decodeTransaction.GasPrice())
		fmt.Println("Value:",decodeTransaction.Value())

		signature,_ := util.SignPayload("0x63949701cd0e1cc04dfea0afbf410968f10ff4b6", decodeTransaction.To().Hex(), decodeTransaction.Data(), decodeTransaction.Gas(), decodeTransaction.Nonce(), "0x3Ca0963A2b3bEeeD3b3EC0d3b6Cae6D99B4855e9")

		fmt.Println("signature:",signature)

		relaySignerService := new(service.RelaySignerService)

		code := relaySignerService.SendMetatransaction(common.HexToAddress("0x63949701cd0e1cc04dfea0afbf410968f10ff4b6"), *decodeTransaction.To(), decodeTransaction.Data(), new(big.Int).SetUint64(decodeTransaction.Gas()), new(big.Int).SetUint64(decodeTransaction.Nonce()), signature, "0x3Ca0963A2b3bEeeD3b3EC0d3b6Cae6D99B4855e9")
		if code == 100{
			log.Println("Failed to send metatransaction")
		}else{
			log.Println("sucess")
		}

		w.Write([]byte("success"))
	}else{
		r.Body=rdr2
		fmt.Println("Is not a rawTransaction")
		
		serveReverseProxy("http://34.75.103.207:4545",w,r)
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

func setupRoutes() {
	fmt.Println("Init RelaySigner")
	http.HandleFunc("/", signTransaction)
	http.ListenAndServe(":9002", nil)
}