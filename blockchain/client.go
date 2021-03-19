package blockchain

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lacchain/gas-relay-signer/model"
	relay "github.com/lacchain/gas-relay-signer/blockchain/contracts"
	"github.com/lacchain/gas-relay-signer/errors"
	log "github.com/lacchain/gas-relay-signer/util"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

//const RelayABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_blocksFrequency\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_accountIngress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSigner\",\"type\":\"bool\"}],\"name\":\"BadSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIRelayHub.ErrorCode\",\"name\":\"errorCode\",\"type\":\"uint8\"}],\"name\":\"BadTransactionSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractDeployed\",\"type\":\"address\"}],\"name\":\"ContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"countExceeded\",\"type\":\"uint8\"}],\"name\":\"GasLimitExceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsedLastBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"averageLastBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGasLimit\",\"type\":\"uint256\"}],\"name\":\"GasLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsedLastBlocks\",\"type\":\"uint256\"}],\"name\":\"GasUsedByTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"NodeBlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"Recalculated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Relayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"TransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relay\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"charge\",\"type\":\"uint256\"}],\"name\":\"TransactionRelayed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNode\",\"type\":\"address\"}],\"name\":\"addNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"}],\"name\":\"deleteNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_byteCode\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"senderSignature\",\"type\":\"bytes\"}],\"name\":\"deployMetaTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"deployedAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasUsedLastBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMsgSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"encodedFunction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"senderSignature\",\"type\":\"bytes\"}],\"name\":\"relayMetaTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accountIngress\",\"type\":\"address\"}],\"name\":\"setAccounIngress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_blocksFrequency\",\"type\":\"uint8\"}],\"name\":\"setBlocksFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGasUsed\",\"type\":\"uint256\"}],\"name\":\"setGasUsedLastBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
const gasUsedByRelay = 300000
const relayMetaTxMethod = "relayMetaTx"
const deployMetaTxMethod = "deployMetaTx"

//Client to manage connection to Ethereum
type Client struct {
	client *ethclient.Client
}

//Connect to Ethereum
func (ec *Client) Connect(nodeURL string) error {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		msg := fmt.Sprintf("Can't connect to node %s", nodeURL)
		err = errors.FailedConnection.Wrapf(err,msg)
		return err
	}

	log.GeneralLogger.Println("Connected to Ethereum Node:", nodeURL)
	ec.client = client
	return nil
}

//Close ethereum connection
func (ec *Client) Close() {
	ec.client.Close()
}

//ConfigTransaction from ethereum address contract
func (ec *Client) ConfigTransaction(key *ecdsa.PrivateKey, gasLimit uint64) (*bind.TransactOpts, error) {	
	auth := bind.NewKeyedTransactor(key)

	nonce, err := ec.client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		msg := fmt.Sprintf("can't get pending nonce for:%s",auth.From)
		err = errors.FailedConfigTransaction.Wrapf(err,msg)
		return nil, err
	}

	gasPrice, err := ec.client.SuggestGasPrice(context.Background())
	if err != nil {
		msg := fmt.Sprintf("can't get gas price suggested")
		err = errors.FailedConfigTransaction.Wrapf(err,msg)
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = gasLimit + gasUsedByRelay // in units
	auth.GasPrice = gasPrice

	log.GeneralLogger.Printf("OptionsTransaction=[From:0x%x,nonce:%d,gasPrice:%s,gasLimit:%d", auth.From,nonce,gasPrice,auth.GasLimit)

	return auth, nil
}

//SendMetatransaction into blockchain
func (ec *Client) SendMetatransaction(contractAddress common.Address, options *bind.TransactOpts, to *common.Address, signingData []byte, v uint8, r [32]byte, s [32]byte ) (*common.Hash,error) {
	contract, err := relay.NewRelay(contractAddress, ec.client)
	if err != nil {
		msg := fmt.Sprintf("can't instance RelayHub contract %s", contractAddress)
		err = errors.FailedContract.Wrapf(err,msg)
		return nil, err
	}

	log.GeneralLogger.Println("RelayHub Contract instanced:",contractAddress.Hex())

	/*log.GeneralLogger.Println("Metatransaction-from:", from.Hex())
	if (to!=nil){
		log.GeneralLogger.Println("Metatransaction to:", to.Hex())
	}
	log.GeneralLogger.Println("Metatransaction encodedFunction:", hexutil.Encode(encodedFunction))
	log.GeneralLogger.Println("Metatransaction gasLimit:", gasLimit)
	log.GeneralLogger.Println("Metatransaction nonce:", nonce)
	log.GeneralLogger.Println("Metatransaction signature:", hexutil.Encode(signature))
*/
	var tx *types.Transaction
	
	if (to != nil){
		tx, err = contract.RelayMetaTx(options, signingData, v, r, s)
	}else{
		tx, err = contract.DeployMetaTx(options, signingData, v, r, s)
	}
	if err != nil {
		msg := fmt.Sprintf("failed executing contract")
		err = errors.BadTransaction.Wrapf(err,msg)
		return nil, err
	}

	if (len(tx.Hash()) == 0){
		msg := fmt.Sprintf("failed executing transaction relay Metatransaction")
		err = errors.FailedTransaction.Wrapf(err,msg)
		return nil, err
	}
	log.GeneralLogger.Printf("MetaTrx sent: %s", tx.Hash().Hex())

	transactionHash := tx.Hash()

	return &transactionHash, nil
}

/*SimulateTransaction DEPRECATED*/
func (ec *Client) SimulateTransaction(nodeURL string, from common.Address, tx *types.Transaction) bool {
	client, err := rpc.DialHTTP(nodeURL)
    if err != nil {
        log.GeneralLogger.Fatal(err)
    }
    defer client.Close()
	
	var result string
	err = client.Call(&result,"eth_call",createCallMsgFromTransaction(from, tx), "latest")
	if err != nil {
		log.GeneralLogger.Fatal("Cannot not get revert reason: " + err.Error())
		return false
	}
	log.GeneralLogger.Println("result:",result)
	value := new(big.Int)
	
	hexResult := strings.Replace(result, "0x", "", -1)
	value.SetString(hexResult, 16)
	log.GeneralLogger.Println("value:",value)
	if value.Int64() == 0 {
		log.GeneralLogger.Println("no error message or out of gas")
		return false
	}
	return true
}

/*DEPRECATED*/
func createCallMsgFromTransaction(from common.Address, tx *types.Transaction) model.CallRequest {
	
	log.GeneralLogger.Println("Call From:",from.Hex())
	log.GeneralLogger.Println("Call To:",tx.To().Hex())
	log.GeneralLogger.Println("Call Data:",hexutil.Encode(tx.Data()))
	log.GeneralLogger.Println("Call GasLimit:",hexutil.EncodeUint64(tx.Gas()))

	return model.CallRequest{
		From: from.Hex(),
		To: tx.To().Hex(),
		Gas: hexutil.EncodeUint64(tx.Gas()),
		Data: hexutil.Encode(tx.Data()),
	}
}

//GenerateTransaction ...DEPRECATED
func (ec *Client)GenerateTransaction(gasLimitTx uint64, relayAddress common.Address, from common.Address, to *common.Address, encodedFunction []byte, gasLimit *big.Int, nonce *big.Int, signature []byte, senderSignature []byte) (*types.Transaction){
	testabi, err := abi.JSON(strings.NewReader(relay.RelayABI))
	if err != nil{
		log.GeneralLogger.Println("Error decoding ABI")
	}

	var bytesData []byte

	if (to!=nil){
		bytesData, _ = testabi.Pack(relayMetaTxMethod,from,*to,encodedFunction,gasLimit,nonce,signature,senderSignature)
	}else{
		bytesData, _ = testabi.Pack(deployMetaTxMethod,from,encodedFunction,gasLimit,nonce,signature,senderSignature)
	}
		tx := types.NewTransaction(0, relayAddress, big.NewInt(0), gasLimitTx, big.NewInt(0), bytesData)
	return tx
}

//GetTransactionReceipt ...
func (ec *Client) GetTransactionReceipt(transactionHash common.Hash)(*types.Receipt, error){
	receipt, err := ec.client.TransactionReceipt(context.Background(), transactionHash)
    if err != nil {
		msg := fmt.Sprintf("failed get transaction receipt %s", transactionHash.Hex())
		err = errors.CallBlockchainFailed.Wrapf(err,msg)
        return nil, err
	}
		
	log.GeneralLogger.Printf("Receipt of Tx:%s was returned with Status:%d in blockNumber:%s",transactionHash.Hex(),receipt.Status,receipt.BlockNumber)

	return receipt,nil;
}

//GetTransactionCount ...
func (ec *Client) GetTransactionCount(contractAddress common.Address, address common.Address)(*big.Int, error){
	contract, err := relay.NewRelay(contractAddress, ec.client)
	if err != nil {
		msg := fmt.Sprintf("can't instance RelayHub contract %s", contractAddress)
		err = errors.FailedContract.Wrapf(err,msg)
		return nil,err
	}

	log.GeneralLogger.Println("RelayHub Contract instanced:",contractAddress.Hex())

	count, err := contract.GetNonce(&bind.CallOpts{}, address)
	
	if err != nil {
		msg := fmt.Sprintf("failed get transaction count for %s", address.Hex())
		err = errors.CallBlockchainFailed.Wrapf(err,msg)
		return nil,err
	}

	return count,nil
}

//DecreaseGasUsed into blockchain
func (ec *Client) DecreaseGasUsed(contractAddress common.Address, options *bind.TransactOpts, gasUsed *big.Int) (*common.Hash,error) {
	contract, err := relay.NewRelay(contractAddress, ec.client)
	if err != nil {
		msg := fmt.Sprintf("can't instance RelayHub contract %s", contractAddress)
		err = errors.FailedContract.Wrapf(err,msg)
		return nil,err
	}

	log.GeneralLogger.Println("RelayHub Contract instanced:",contractAddress.Hex())
	
	var tx *types.Transaction
	
	tx, err = contract.IncreaseGasUsed(options, gasUsed)
	
	if err != nil {
		msg := fmt.Sprintf("failed executing contract")
		err = errors.BadTransaction.Wrapf(err,msg)
		return nil,err
	}

	if (len(tx.Hash()) == 0){
		msg := fmt.Sprintf("failed executing transaction decrease gas")
		err = errors.FailedTransaction.Wrapf(err,msg)
		return nil,err
	}
	log.GeneralLogger.Printf("Decreased Gas Tx sent: %s", tx.Hash().Hex())

	transactionHash := tx.Hash()

	return &transactionHash, nil
}

//IsTargetPermitted DEPRECATED...
func (ec *Client) IsTargetPermitted(contractAddress common.Address, address common.Address)(bool, error){
	contract, err := relay.NewAccount(contractAddress, ec.client)
	if err != nil {
		msg := fmt.Sprintf("can't instance Account Permissioning contract %s", contractAddress)
		err = errors.FailedContract.Wrapf(err,msg)
		return false,err
	}

	log.GeneralLogger.Println("Account Permissioning Contract instanced:",contractAddress.Hex())

	isPermitted, err := contract.DestinationPermitted(&bind.CallOpts{}, address)

	if err != nil {
		return false,err
	}

	return isPermitted, nil
}