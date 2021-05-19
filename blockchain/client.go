package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	relay "github.com/lacchain/gas-relay-signer/blockchain/contracts"
	"github.com/lacchain/gas-relay-signer/errors"
	"github.com/lacchain/gas-relay-signer/model"
	log "github.com/lacchain/gas-relay-signer/audit"
)

const (
	gasUsedByRelay = 300000
	relayMetaTxMethod = "relayMetaTx"
	deployMetaTxMethod = "deployMetaTx"
)

//Client to manage connection to Ethereum
type Client struct {
	client *ethclient.Client
}

//Connect to Ethereum
func (ec *Client) Connect(nodeURL string) error {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		msg := fmt.Sprintf("Can't connect to node %s", nodeURL)
		err = errors.FailedConnection.Wrapf(err, msg, -32100)
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
		msg := fmt.Sprintf("can't get pending nonce for:%s", auth.From)
		err = errors.FailedConfigTransaction.Wrapf(err, msg, -32603)
		return nil, err
	}

	gasPrice, err := ec.client.SuggestGasPrice(context.Background())
	if err != nil {
		msg := "can't get gas price suggested"
		err = errors.FailedConfigTransaction.Wrapf(err, msg, -32603)
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)                // in wei
	auth.GasLimit = gasLimit + gasUsedByRelay // in units
	auth.GasPrice = gasPrice

	log.GeneralLogger.Printf("OptionsTransaction=[From:0x%x,nonce:%d,gasPrice:%s,gasLimit:%d", auth.From, nonce, gasPrice, auth.GasLimit)

	return auth, nil
}

//SendMetatransaction into blockchain
func (ec *Client) SendMetatransaction(contractAddress common.Address, options *bind.TransactOpts, to *common.Address, signingData []byte, v uint8, r [32]byte, s [32]byte) (*common.Hash, error) {
	contract, err := relay.NewRelay(contractAddress, ec.client)
	if err != nil {
		msg := fmt.Sprintf("can't instance RelayHub contract %s", contractAddress)
		err = errors.FailedContract.Wrapf(err, msg, -32603)
		return nil, err
	}

	log.GeneralLogger.Println("RelayHub Contract instanced:", contractAddress.Hex())
	log.GeneralLogger.Println("Metatransaction signingData:", hexutil.Encode(signingData))
	
	var tx *types.Transaction

	if to != nil {
		tx, err = contract.RelayMetaTx(options, signingData, v, r, s)
	} else {
		tx, err = contract.DeployMetaTx(options, signingData, v, r, s)
	}
	if err != nil {
		msg := "failed executing smart contract"
		err = errors.FailedTransaction.Wrapf(err, msg, -32603)
		return nil, err
	}

	if len(tx.Hash()) == 0 {
		msg := "failed executing transaction relay Metatransaction"
		err = errors.FailedTransaction.Wrapf(err, msg, -32603)
		return nil, err
	}
	log.GeneralLogger.Printf("MetaTransaction sent: %s", tx.Hash().Hex())

	transactionHash := tx.Hash()

	return &transactionHash, nil
}

func (ec *Client) SimulateTransaction(nodeURL string, from common.Address, tx *types.Transaction) (uint64,error) {
	client, err := rpc.DialHTTP(nodeURL)
	if err != nil {
		msg := fmt.Sprintf("Can't connect to node %s", nodeURL)
		err = errors.FailedConnection.Wrapf(err, msg, -32100)
		return 7,err
	}
	defer client.Close()

	var result string
	err = client.Call(&result, "eth_call", createCallMsgFromTransaction(from, tx), "pending")
	if err != nil {
		msg := "Can not get revert reason"
		err = errors.FailedTransaction.Wrapf(err, msg, -32603)
		return 7, err
	}
	value := new(big.Int)
    log.GeneralLogger.Println("simulate result:",result)
	hexResult := strings.Replace(result, "0x", "", -1)
	value.SetString(string([]byte(hexResult)[:64]), 16)
	log.GeneralLogger.Println("simulate transaction result code:", value)
	
	return value.Uint64(),nil
}

func createCallMsgFromTransaction(from common.Address, tx *types.Transaction) model.CallRequest {
	log.GeneralLogger.Printf("Call=[From:%s,To:%s,Data:%s,gasLimit:%s", from.Hex(), tx.To().Hex(), hexutil.Encode(tx.Data()), hexutil.EncodeUint64(tx.Gas()))

	return model.CallRequest{
		From: from.Hex(),
		To:   tx.To().Hex(),
		Gas:  hexutil.EncodeUint64(tx.Gas()),
		Data: hexutil.Encode(tx.Data()),
	}
}

func (ec *Client) GenerateTransaction(options *bind.TransactOpts, to *common.Address, relayAddress common.Address, signingData []byte, v uint8, r , s [32]byte) (*types.Transaction,error) {
	testabi, err := abi.JSON(strings.NewReader(relay.RelayABI))
	if err != nil {
		msg := "Error decoding ABI"
		err = errors.FailedTransaction.Wrapf(err, msg, -32603)
		return nil, err
	}

	var bytesData []byte

	if to!=nil {
		bytesData, err = testabi.Pack(relayMetaTxMethod, signingData, v, r, s)
	} else {
		fmt.Println("DEPLOYYYY")
		bytesData, err = testabi.Pack(deployMetaTxMethod, signingData, v, r, s)
	}

	if err != nil {
		msg := "Error encoding transaction"
		err = errors.FailedTransaction.Wrapf(err, msg, -32603)
		return nil, err
	}

	tx := types.NewTransaction(options.Nonce.Uint64(), relayAddress, big.NewInt(0), options.GasLimit, big.NewInt(0), bytesData)
	return tx,nil
}

//GetTransactionReceipt ...
func (ec *Client) GetTransactionReceipt(transactionHash common.Hash) (*types.Receipt, error) {
	receipt, err := ec.client.TransactionReceipt(context.Background(), transactionHash)
	if err != nil {
		msg := fmt.Sprintf("failed get transaction receipt %s", transactionHash.Hex())
		err = errors.CallBlockchainFailed.Wrapf(err, msg, -32603)
		return nil, err
	}

	log.GeneralLogger.Printf("Receipt of Tx:%s was returned with Status:%d in blockNumber:%s", transactionHash.Hex(), receipt.Status, receipt.BlockNumber)

	return receipt, nil
}

//GetTransactionCount ...
func (ec *Client) GetTransactionCount(contractAddress common.Address, address common.Address) (*big.Int, error) {
	contract, err := relay.NewRelay(contractAddress, ec.client)
	if err != nil {
		msg := fmt.Sprintf("can't instance RelayHub contract %s", contractAddress)
		err = errors.FailedContract.Wrapf(err, msg, -32603)
		return nil, err
	}

	log.GeneralLogger.Println("RelayHub Contract instanced:", contractAddress.Hex())

	count, err := contract.GetNonce(&bind.CallOpts{}, address)

	if err != nil {
		msg := fmt.Sprintf("failed get transaction count for %s", address.Hex())
		err = errors.CallBlockchainFailed.Wrapf(err, msg, -32603)
		return nil, err
	}

	return count, nil
}

//DecreaseGasUsed into blockchain
func (ec *Client) DecreaseGasUsed(contractAddress common.Address, options *bind.TransactOpts, gasUsed *big.Int) (*common.Hash, error) {
	contract, err := relay.NewRelay(contractAddress, ec.client)
	if err != nil {
		msg := fmt.Sprintf("can't instance RelayHub contract %s", contractAddress)
		err = errors.FailedContract.Wrapf(err, msg, -32603)
		return nil, err
	}

	log.GeneralLogger.Println("RelayHub Contract instanced:", contractAddress.Hex())

	var tx *types.Transaction

	tx, err = contract.IncreaseGasUsed(options, gasUsed)

	if err != nil {
		msg := "failed executing contract"
		err = errors.FailedTransaction.Wrapf(err, msg, -32603)
		return nil, err
	}

	if len(tx.Hash()) == 0 {
		msg := "failed executing transaction decrease gas"
		err = errors.FailedTransaction.Wrapf(err, msg, -32603)
		return nil, err
	}
	log.GeneralLogger.Printf("Decreased Gas Tx sent: %s", tx.Hash().Hex())

	transactionHash := tx.Hash()

	return &transactionHash, nil
}

//GetBlockByNumber ...
func (ec *Client) GetBlockByNumber(contractAddress common.Address, blockNumber *big.Int) (*types.Block, uint64, error){
	response, err := ec.client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		msg := fmt.Sprintf("failed get block by number %d", blockNumber.Uint64())
		err = errors.CallBlockchainFailed.Wrapf(err, msg, -32603)
		return nil, 0, err
	}

	gasLimit, err := ec.GetGasLimit(contractAddress)
	if err != nil{
		return nil, 0, err
	}

	return response, gasLimit.Uint64(), nil
}

//GetGasLimit ...
func (ec *Client) GetGasLimit(contractAddress common.Address)(*big.Int,error){
	contract, err := relay.NewRelay(contractAddress, ec.client)
	if err != nil {
		msg := fmt.Sprintf("can't instance RelayHub contract %s", contractAddress)
		err = errors.FailedContract.Wrapf(err, msg, -32603)
		return nil, err
	}

	log.GeneralLogger.Println("RelayHub Contract instanced:", contractAddress.Hex())

	gasLimit, err := contract.GetMaxGasBlockLimit(&bind.CallOpts{})

	if err != nil {
		msg := fmt.Sprintf("failed get gasLimit from %s",contractAddress.Hex())
		err = errors.CallBlockchainFailed.Wrapf(err, msg, -32603)
		return nil, err
	}

	return gasLimit, nil
}