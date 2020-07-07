package blockchain

import (
	"context"
	"time"
	"fmt"
	"log"
	"math/big"
//	"strconv"
	"strings"
//	"os"
	"crypto/ecdsa"
//	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lacchain/gas-relay-signer/model"
	relay "github.com/lacchain/gas-relay-signer/blockchain/contracts"
	"github.com/lacchain/gas-relay-signer/errors"
	l "github.com/lacchain/gas-relay-signer/util"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const RelayABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"GasLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"name\":\"GasUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"Hashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Relayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relay\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"charge\",\"type\":\"uint256\"}],\"name\":\"TransactionRelayed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"encodedFunction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"relayMetaTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"setGasLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

//Client to manage Connection to Ethereum
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

	l.GeneralLogger.Println("Connected to Ethereum Node:", nodeURL)
	ec.client = client
	return nil
}

//Close ethereum connection
func (ec *Client) Close() {
	ec.client.Close()
}

//ConfigTransaction from ethereum address contract
func (ec *Client) ConfigTransaction(key *ecdsa.PrivateKey, gasLimit uint64) (*bind.TransactOpts, error) {
/*	keystore, err := os.Open(keyStorePath)
	defer keystore.Close()
    if err != nil {
		msg := fmt.Sprintf("could not load keystore from location %s",keyStorePath)
		err = errors.FailedKeystore.Wrapf(err,msg)
		return nil,err
    }
*/	
	auth := bind.NewKeyedTransactor(key)

	nonce, err := ec.client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		msg := fmt.Sprintf("can't get pending nonce for:",auth.From)
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
	auth.GasLimit = gasLimit + 150000 // in units
	auth.GasPrice = gasPrice

	l.GeneralLogger.Printf("OptionsTransaction=[From:0x%x,nonce:%d,gasPrice:%s,gasLimit:%d", auth.From,nonce,gasPrice,auth.GasLimit)

	return auth, nil
}

//SendMetatransaction into blockchain
func (ec *Client) SendMetatransaction(contractAddress common.Address, options *bind.TransactOpts, from common.Address, to common.Address, encodedFunction []byte, gasLimit *big.Int, nonce *big.Int, signature []byte) (error, *common.Hash) {
	contract, err := relay.NewRelay(contractAddress, ec.client)
	if err != nil {
		msg := fmt.Sprintf("can't instance RelayHub contract %s", contractAddress)
		err = errors.FailedContract.Wrapf(err,msg)
		return err, nil
	}

	log.Println("RelayHub Contract instanced:",contractAddress.Hex())

	log.Println("from:", from.Hex())
	log.Println("to:", to.Hex())
	log.Println("encodedFunction:", hexutil.Encode(encodedFunction))
	log.Println("gasLimit:", gasLimit)
	log.Println("nonce:", nonce)
	log.Println("signature:", hexutil.Encode(signature))

	tx, err := contract.RelayMetaTx(options, from, to, encodedFunction, gasLimit, nonce, signature)
	if err != nil {
		msg := fmt.Sprintf("failed executing contract")
		err = errors.BadTransaction.Wrapf(err,msg)
		return err, nil
	}

	if (len(tx.Hash()) == 0){
		msg := fmt.Sprintf("failed execute transaction")
		err = errors.FailedTransaction.Wrapf(err,msg)
		return err, nil
	}
	log.Printf("Tx sent: %s", tx.Hash().Hex())

	transactionHash := tx.Hash()

	return nil, &transactionHash
}

func (ec *Client) SimulateTransaction(nodeURL string, from common.Address, tx *types.Transaction) bool {
	client, err := rpc.DialHTTP(nodeURL)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()
	
	var result string
	err = client.Call(&result,"eth_call",createCallMsgFromTransaction(from, tx), "latest")
	if err != nil {
		log.Fatal("Cannot not get revert reason: " + err.Error())
		return false
	}
	fmt.Println("result:",result)
	value := new(big.Int)
	
	hexResult := strings.Replace(result, "0x", "", -1)
	value.SetString(hexResult, 16)
	fmt.Println("value:",value)
	if value.Int64() == 0 {
		fmt.Println("no error message or out of gas")
		return false
	}
	return true
}

func createCallMsgFromTransaction(from common.Address, tx *types.Transaction) model.CallRequest {
	
	fmt.Println("Call From:",from.Hex())
	fmt.Println("Call To:",tx.To().Hex())
	fmt.Println("Call Data:",hexutil.Encode(tx.Data()))
	fmt.Println("Call GasLimit:",hexutil.EncodeUint64(tx.Gas()))

	return model.CallRequest{
		From: from.Hex(),
		To: tx.To().Hex(),
		Gas: hexutil.EncodeUint64(tx.Gas()),
		Data: hexutil.Encode(tx.Data()),
	}
}

//GenerateTransaction ...
func (ec *Client)GenerateTransaction(gasLimitTx uint64, relayAddress common.Address, from common.Address, to common.Address, encodedFunction []byte, gasLimit *big.Int, nonce *big.Int, signature []byte) (*types.Transaction){
	testabi, err := abi.JSON(strings.NewReader(RelayABI))
	if err != nil{
		fmt.Println("Error decoding ABI")
	}
	bytesData, _ := testabi.Pack("relayMetaTx",from,to,encodedFunction,gasLimit,nonce,signature)
	tx := types.NewTransaction(100, relayAddress, big.NewInt(0), gasLimitTx, big.NewInt(0), bytesData)
	return tx
}

//GetTransactionReceipt ...
func (ec *Client) GetTransactionReceipt(transactionHash common.Hash)(*big.Int, string, error){
	receipt, err := ec.client.TransactionReceipt(context.Background(), transactionHash)
        if err != nil {
            log.Fatal(err)
		}
		
	log.Println("Status:",receipt.Status)
	log.Println("BlockNumber:",receipt.BlockNumber)

	block, err := ec.client.BlockByNumber(context.Background(), receipt.BlockNumber)
    if err != nil {
        log.Fatal(err)
	}
	
	log.Println("block time:",block.Time())

	ts := time.Unix(int64(block.Time()),0).UTC()

	return receipt.BlockNumber, ts.String(), nil
}