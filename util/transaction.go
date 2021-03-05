package util

import(
	"fmt"
//	"errors"
	"log"
	"strconv"
	"encoding/hex"
	"crypto/ecdsa"
	sha "golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/lacchain/gas-relay-signer/sha3"
	"github.com/lacchain/gas-relay-signer/errors"
)

//SignPayload ...
func SignPayload(_privateKey, signingAddr string, destinationAddress *common.Address, encodedFunction []byte, gasLimit, nonce uint64) ([]byte, error){
	log.Println("rawData", hexutil.Encode(encodedFunction))

	//nonce := await txRelay.getNonce.call(signingAddr)
	d := sha.NewLegacyKeccak256()
	d.Write(encodedFunction)

	//Tight packing, as Solidity sha3 does
	hash := Hash(signingAddr,destinationAddress,d.Sum(nil),strconv.FormatUint(gasLimit, 10),strconv.FormatUint(nonce, 10))
	
	fmt.Println("messageHashed:",hexutil.Encode(hash))

	privateKey, err := crypto.HexToECDSA(_privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
	}
	
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Node Address:",address)

	if err != nil {
        log.Fatal(err)
    }

	hash2 := sha3.SoliditySHA3(
		sha3.String("\x19Ethereum Signed Message:\n32"),
		sha3.Bytes32(hash),
	)

	fmt.Println("messageHashed2:",hexutil.Encode(hash2))

	sig, err := Sign(hash2, privateKey)
	if err != nil {
        log.Fatal(err)
    }
						
	return sig,nil 
}

//Hash ...
func Hash(from string, to *common.Address, encodedFunction []byte, gasLimit, nonce string) ([]byte) {
	var hash []byte
	if (to != nil){
		hash = sha3.SoliditySHA3(
			// types
			[]string{"address", "address", "bytes32", "uint256", "uint256"},

			// values
			[]interface{}{
				from,
				to.Hex(),
				encodedFunction,
				gasLimit,
				nonce,
			},
		)
	}else{
		hash = sha3.SoliditySHA3(
			// types
			[]string{"address", "bytes32", "uint256", "uint256"},

			// values
			[]interface{}{
				from,
				encodedFunction,
				gasLimit,
				nonce,
			},
		)
	}
	
	return hash
}

//Sign ...
func Sign(hash []byte, privateKey *ecdsa.PrivateKey) ([]byte,error) {
	signature, err := crypto.Sign(hash, privateKey)
    if err != nil {
        log.Fatal(err)
	}

	fmt.Println("signature length:",len(signature))

	if len(signature) != 65 {
		return nil,errors.New("signature malformed")
	}

	fmt.Println("V byte:",hexutil.Encode(signature[64:65]))
	var v1 string
	if (hexutil.Encode(signature[64:65]) == "0x00") {
		v1 = "1b"
	} else if (hexutil.Encode(signature[64:65]) == "0x01") {
		v1 = "1c"
	} else {
		return nil, errors.New("bad V parameter")
	}

	vParameter,err := hex.DecodeString(v1)
	if err != nil {
        log.Println(err)
	}

	fmt.Println("V Parameter:",vParameter)

	signature[64] = vParameter[0]

	fmt.Println("signature:",hexutil.Encode(signature))

	return signature,nil
}

//GetTransaction ...
func GetTransaction(rawTx string)(*types.Transaction, error){
	rawTxBytes, err := hex.DecodeString(rawTx)

	if err != nil {
		msg := fmt.Sprintf("Error Decoding Raw Transaction")
		err = errors.MalformedRawTransaction.Wrapf(err,msg)
		return nil, err
    }

    tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)
	
	return tx,nil
}



