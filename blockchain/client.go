package blockchain

import (
	"context"
	"time"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	relay "github.com/lacchain/gas-relay-signer/blockchain/contracts"
	"github.com/lacchain/gas-relay-signer/errors"
	l "github.com/lacchain/gas-relay-signer/util"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

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
func (ec *Client) ConfigTransaction(keyStorePath,keystorepass string) (*bind.TransactOpts, error) {
	keystore, err := os.Open(keyStorePath)
	defer keystore.Close()
    if err != nil {
		msg := fmt.Sprintf("could not load keystore from location %s",keyStorePath)
		err = errors.FailedKeystore.Wrapf(err,msg)
		return nil,err
    }
	
	auth, err := bind.NewTransactor(keystore, keystorepass)
	if err != nil {
		msg := fmt.Sprintf("can't open the keystore")
		err = errors.FailedKeystore.Wrapf(err,msg)
		return nil, err
	}

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
	auth.GasLimit = uint64(200000) // in units
	auth.GasPrice = gasPrice

	l.GeneralLogger.Printf("OptionsTransaction=[From:0x%x,nonce:%d,gasPrice:%s,gasLimit:%d,gas:%s", auth.From,nonce,gasPrice,auth.GasLimit,0)

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