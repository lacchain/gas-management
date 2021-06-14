/*
	RelaySigner Service
	version 0.9
	author: Adrian Pareja Abarca
	email: adriancc5.5@gmail.com
*/

package service

import (
	"encoding/json"
//	"io/ioutil"
	"os"
	"fmt"
	"strings"
	"math/big"
	"encoding/hex"
	"crypto/ecdsa"
	sha "golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/lacchain/gas-relay-signer/rpc"
	bl "github.com/lacchain/gas-relay-signer/blockchain"
	log "github.com/lacchain/gas-relay-signer/audit"
	"github.com/lacchain/gas-relay-signer/model"
	"github.com/lacchain/gas-relay-signer/errors"
)

const RelayABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_blocksFrequency\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_accountIngress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"AccountIngressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIRelayHub.ErrorCode\",\"name\":\"errorCode\",\"type\":\"uint8\"}],\"name\":\"BadTransactionSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"blocksFrequency\",\"type\":\"uint8\"}],\"name\":\"BlockFrequencyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relay\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractDeployed\",\"type\":\"address\"}],\"name\":\"ContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"countExceeded\",\"type\":\"uint8\"}],\"name\":\"GasLimitExceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsedLastBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"averageLastBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGasLimit\",\"type\":\"uint256\"}],\"name\":\"GasLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsedLastBlocks\",\"type\":\"uint256\"}],\"name\":\"GasUsedByTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsedRelayHub\",\"type\":\"uint256\"}],\"name\":\"GasUsedRelayHubChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxGasBlockLimit\",\"type\":\"uint256\"}],\"name\":\"MaxGasBlockLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newNode\",\"type\":\"address\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"NodeBlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldNode\",\"type\":\"address\"}],\"name\":\"NodeDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"decodedFunction\",\"type\":\"bytes\"}],\"name\":\"Parameters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"Recalculated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Relayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relay\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"name\":\"TransactionRelayed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNode\",\"type\":\"address\"}],\"name\":\"addNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"deleteNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasUsedLastBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accountIngress\",\"type\":\"address\"}],\"name\":\"setAccounIngress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_blocksFrequency\",\"type\":\"uint8\"}],\"name\":\"setBlocksFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGasUsed\",\"type\":\"uint256\"}],\"name\":\"setGasUsedLastBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasUsedRelayHub\",\"type\":\"uint256\"}],\"name\":\"setGasUsedRelayHub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxGasBlockLimit\",\"type\":\"uint256\"}],\"name\":\"setMaxGasBlockLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signingData\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"relayMetaTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signingData\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"deployMetaTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"deployedAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMsgSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"name\":\"increaseGasUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

const ENVIRONMENT_KEY_NAME = "WRITER_KEY"

//RelaySignerService is the main service
type RelaySignerService struct {
	// The service's configuration
	Config *model.Config
}

//Init configuration parameters
func (service *RelaySignerService) Init(_config *model.Config)(error){
	service.Config = _config

	for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }

    key, exist := os.LookupEnv(ENVIRONMENT_KEY_NAME)
    if !exist {
		return errors.FailedReadEnv.New("Environment variable WRITER_KEY not set", -32602)
    } 

	privateKey, err := crypto.HexToECDSA(string(key[2:66]))
    if err != nil {
        return errors.FailedKeyConfig.New("Invalid ECDSA Key", -32602)
	}

	publicKey := privateKey.Public()
	_, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        return errors.FailedKeyConfig.New("Invalid ECDSA Public Key", -32602)
	}

	service.Config.Application.Key = string(key[2:66])

	if service.Config.Security.PermissionsEnabled{
		if !(common.IsHexAddress(service.Config.Security.AccountContractAddress)){
			return errors.InvalidAddress.New("Invalid Account Smart Contract Address",-32608)
		}
	}
	return nil
}

//SendMetatransaction to blockchain
func (service *RelaySignerService) SendMetatransaction(id json.RawMessage, to *common.Address, gasLimit uint64, signingData []byte, v uint8, r, s [32]byte) (*rpc.JsonrpcMessage) {
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.NodeURL)
	if err != nil {
		HandleError(id, err)
	}
	defer client.Close()
	
	privateKey, err := crypto.HexToECDSA(service.Config.Application.Key)
    if err != nil {
        HandleError(id, err)
	}

	/*publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
		err := errors.New("error casting public key to ECDSA", -32602)
		HandleError(id, err)
	}*/
	
	//address := crypto.PubkeyToAddress(*publicKeyECDSA)

	/*options, err := client.ConfigTransaction(privateKey,gasLimit,true)
	if err != nil {
		return HandleError(id, err)
	}*/
	contractAddress := common.HexToAddress(service.Config.Application.ContractAddress)

	/*callTx,err := client.GenerateTransaction(options,to,contractAddress,signingData,v,r,s)
	if err != nil {
		return HandleError(id, err)
	}
	codeResponse,err := client.SimulateTransaction(service.Config.Application.NodeURL,address,callTx)
	if err != nil {
		return HandleError(id, err)
	}

	if( codeResponse != 8){
		log.GeneralLogger.Println("Transaction will fail, then is rejected")
		switch codeResponse {
		case 0:
			err = errors.New("transaction gas limit exceeds block gas limit",-32005) 
		case 1,6:
			err = errors.New("Invalid Signature",-32002) 
		case 2:
			err = errors.New("Incorrect Nonce",-32006) 
		case 3:
			err = errors.New("Intrinsic gas exceeds gas limit",-32003) 
		case 4:
			err = errors.New("Recepient is not a contract",-32004) 
		case 5:
			err = errors.New("Empty code can't be deployed",-32001) 
		case 7:
			err = errors.New("Invalid recipient smart contract ",-32007) 
		}
		
		return HandleError(id, err)
	}*/

	optionsSendTransaction, err := client.ConfigTransaction(privateKey,gasLimit,false)
	if err != nil {
		return HandleError(id, err)
	}
	tx, err := client.SendMetatransaction(contractAddress, optionsSendTransaction, to, signingData, v, r, s)
	if err != nil {
		return HandleError(id, err)
	}

	log.GeneralLogger.Println("transaction",tx)

	result := new(rpc.JsonrpcMessage)

	result.ID = id
	return result.Response(tx)
}

//GetTransactionReceipt from blockchain
func (service *RelaySignerService) GetTransactionReceipt(id json.RawMessage,transactionID string) (*rpc.JsonrpcMessage){
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.NodeURL)
	if err != nil {
		HandleError(id,err)
	}
	defer client.Close()

	receipt,err := client.GetTransactionReceipt(common.HexToHash(transactionID))
	if err != nil {
		HandleError(id,err)
	}

	var receiptReverted map[string]interface{}

	if receipt!=nil{
		d := sha.NewLegacyKeccak256()
		e := sha.NewLegacyKeccak256()
		d.Write([]byte("ContractDeployed(address,address,address)"))

		eventContractDeployed := hex.EncodeToString(d.Sum(nil))

		e.Write([]byte("TransactionRelayed(address,address,address,bool,bytes)"))
		eventTransactionRelayed := hex.EncodeToString(e.Sum(nil))

		fmt.Println("deployed contract eventKeccak:",eventContractDeployed)
		fmt.Println("transaction relayed eventKeccak:",eventTransactionRelayed)

		for _,log := range receipt.Logs {
		//	fmt.Println(log.Topics[0].Hex())
			if log.Topics[0].Hex() == "0x"+eventContractDeployed {
			//	fmt.Println(hex.EncodeToString(log.Data))
			//	fmt.Println(common.BytesToAddress(log.Data).Hex())
				receipt.ContractAddress = common.BytesToAddress(log.Data)
			}
			if log.Topics[0].Hex() == "0x"+eventTransactionRelayed {
				executed, output := transactionRelayedFailed(id,log.Data)
				if !executed{
					receipt.Status = uint64(0)
					fmt.Println("Reverse Error:",hexutil.Encode(output))

					jsonReceipt,err := json.Marshal(receipt)
					if err != nil{
						HandleError(id,err)
					}

					json.Unmarshal(jsonReceipt, &receiptReverted)
					receiptReverted["revertReason"] = hexutil.Encode(output)
				}
			}
		}
	}
	result := new(rpc.JsonrpcMessage)

	result.ID = id
	if (receiptReverted != nil){
		return result.Response(receiptReverted)
	}
	return result.Response(receipt)
	
}

//GetTransactionCount of account
func (service *RelaySignerService) GetTransactionCount(id json.RawMessage,from string) (*rpc.JsonrpcMessage){
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.NodeURL)
	if err != nil {
		HandleError(id,err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress(service.Config.Application.ContractAddress)
	address := common.HexToAddress(from)

	count,err := client.GetTransactionCount(contractAddress, address)
	if err != nil {
		HandleError(id,err)
	}

	result := new(rpc.JsonrpcMessage)

	result.ID = id
	return result.Response(count)
}

//GetGasLimit of account
func (service *RelaySignerService) GetGasLimit(id json.RawMessage) (uint64){
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.NodeURL)
	if err != nil {
		HandleError(id,err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress(service.Config.Application.ContractAddress)

	gasLimit,err := client.GetMaxBlockGasLimit(contractAddress)
	if err != nil {
		HandleError(id,err)
	}

	log.GeneralLogger.Println("block logic gasLimit:",gasLimit.Uint64())

	return gasLimit.Uint64()
}

//GetBlockByNumber ...
func (service *RelaySignerService) GetBlockByNumber(id json.RawMessage,blockNumber *big.Int) (*rpc.JsonrpcMessage){
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.NodeURL)
	if err != nil {
		HandleError(id,err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress(service.Config.Application.ContractAddress)

	blockDetails,count,err := client.GetBlockByNumber(contractAddress, blockNumber)
	if err != nil {
		HandleError(id,err)
	}

	var blockMap map[string]interface{}
	jsonBlock,err := json.Marshal(blockDetails)
	if err != nil{
		HandleError(id,err)
	}

	json.Unmarshal(jsonBlock, &blockMap)
	blockMap["gasLimit"] = hexutil.EncodeUint64(count)

	result := new(rpc.JsonrpcMessage)

	result.ID = id
	return result.Response(blockMap)
}

//VerifyGasLimit sent a transaction
func (service *RelaySignerService) VerifyGasLimit(gasLimit uint64, id json.RawMessage) (bool){
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.NodeURL)
	if err != nil {
		HandleError(id,err)
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(service.Config.Application.Key)
    if err != nil {
        HandleError(id, err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
		err := errors.New("error casting public key to ECDSA", -32602)
		HandleError(id, err)
	}
	
	nodeAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	contractAddress := common.HexToAddress(service.Config.Application.ContractAddress)

	maxBlockGasLimit,err := client.GetMaxBlockGasLimit(contractAddress)
	if err != nil {
		HandleError(id,err)
	}

	log.GeneralLogger.Println("logic block gasLimit:",maxBlockGasLimit.Uint64())

	nodeGasLimit,err := client.GetGasLimit(contractAddress, nodeAddress)
	if err != nil {
		HandleError(id,err)
	}

	log.GeneralLogger.Println("node gasLimit:",nodeGasLimit.Uint64())

	if (gasLimit > maxBlockGasLimit.Uint64()) || (gasLimit > nodeGasLimit.Uint64()) {
		return false
	}

	return true
}

//VerifySender sent a transaction
func (service *RelaySignerService) VerifySender(sender common.Address, id json.RawMessage) (bool){
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.NodeURL)
	if err != nil {
		HandleError(id,err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress(service.Config.Security.AccountContractAddress)

	isPermitted,err := client.AccountPermitted(contractAddress, sender)
	if err != nil {
		HandleError(id,err)
	}

	log.GeneralLogger.Println("sender is permitted:",isPermitted)

	return isPermitted
}

//DecreaseGasUsed by node
func (service *RelaySignerService) DecreaseGasUsed(id json.RawMessage) (bool){
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.NodeURL)
	if err != nil {
		HandleError(id,err)
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(service.Config.Application.Key)
    if err != nil {
        log.GeneralLogger.Fatal(err)
	}

	options, err := client.ConfigTransaction(privateKey,30000,false)
	if err != nil {
	 	HandleError(id,err)
	}

	contractAddress := common.HexToAddress(service.Config.Application.ContractAddress)

	_, err = client.DecreaseGasUsed(contractAddress, options, new(big.Int).SetUint64(25000))
	if err != nil {
		HandleError(id,err)
	}
	
	return true
}

func transactionRelayedFailed(id json.RawMessage,data []byte)(bool, []byte){
	var transactionRelayedEvent struct {
		Relay   	common.Address
		From      	common.Address
		To        	common.Address
		Executed    bool
		Output	  	[]byte 
	}

	relayHubAbi, err := abi.JSON(strings.NewReader(RelayABI))
	if err != nil {
		HandleError(id,err)
	}

	err = relayHubAbi.Unpack(&transactionRelayedEvent, "TransactionRelayed", data)

	if err != nil {
		fmt.Println("Failed to unpack")
	}

	return transactionRelayedEvent.Executed, transactionRelayedEvent.Output;
}

//HandleError
func HandleError(id json.RawMessage, err error)(*rpc.JsonrpcMessage){
	log.GeneralLogger.Println(err.Error())
	result := new(rpc.JsonrpcMessage)
	result.ID = id
	return result.ErrorResponse(err)
}