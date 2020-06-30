/*
	RelaySigner Service
	version 0.9
	author: Adrian Pareja Abarca
	email: adriancc5.5@gmail.com
*/

package service

import (
	"log"
	"fmt"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	bl "github.com/lacchain/gas-relay-signer/blockchain"
	"github.com/lacchain/gas-relay-signer/model"
	"github.com/lacchain/gas-relay-signer/errors"
)

//RelaySignerService is the main service
type RelaySignerService struct {
	// The service's configuration
	Config *model.Config
}

//Init configuration parameters
func (service *RelaySignerService) Init(_config *model.Config){
	service.Config = _config
}

//SendMetatransaction saving the hash into blockchain
func (service *RelaySignerService) SendMetatransaction(from common.Address, to common.Address, encodedFunction []byte, gasLimit, nonce *big.Int, signature []byte, relayHubContract string) (int) {
	client := new(bl.Client)
	err := client.Connect("http://34.75.103.207:4545")
	if err != nil {
		handleError(err)
	}
	defer client.Close()
	
	options, err := client.ConfigTransaction("/home/adrian/.ethereum/keystore/UTC--2020-06-18T05-39-57.250267184Z--63949701cd0e1cc04dfea0afbf410968f10ff4b6","Peru123.")
	if err != nil {
		return handleError(err)
	}
	contractAddress := common.HexToAddress(relayHubContract)

	err, tx := client.SendMetatransaction(contractAddress, options, from, to, encodedFunction, gasLimit, nonce, signature)
	if err != nil {
		return handleError(err)
	}

	fmt.Println("tx",tx)

	return 200
}

func handleError(err error)(int){
	errorType := errors.GetType(err)
   	switch errorType {
    	case errors.FailedConnection: 
			  log.Fatal(err.Error())
		case errors.FailedKeystore:
			  log.Fatal(err.Error())	
		case errors.FailedConfigTransaction:
			log.Println(err.Error())
			return 100  
		default: 
      		log.Println("otro error:",err)
	   }
	  return 100
}