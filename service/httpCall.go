package service

import(
	"fmt"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/lacchain/gas-relay-signer/audit"
	"github.com/lacchain/gas-relay-signer/rpc"
)

func isPoolEmpty(rpcURL string, id json.RawMessage) (bool,error){
	data := fmt.Sprintf(`{"jsonrpc":"2.0","method":"txpool_besuTransactions",
	"params":[], "id":"%s"}`,id)

	requestBody := []byte(data)

	timeout := time.Duration(5*time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", rpcURL, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type","application/json")

	if err != nil {
		return false,err
	}

	response, err := client.Do(request)
	if err != nil {
		return false,err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return false,err
	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(body))

	var rpcMessage rpc.JsonrpcMessage

	err = json.NewDecoder(rdr1).Decode(&rpcMessage)
	if err != nil {
		return false,err
	}

	audit.GeneralLogger.Println("Transactions in pool:",rpcMessage.String())

	var v []string    
	err = json.Unmarshal(rpcMessage.Result, &v)    
	if err != nil {
		return false,err
	}
	
	if (len(v)>0){
		return false,nil
	}

	return true,nil
}