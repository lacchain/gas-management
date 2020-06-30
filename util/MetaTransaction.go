package util

type MetaTransaction struct {
	From string
	To string
	EncodedFunction []byte
	GasLimit int
	Nonce int
	Signature []byte
}