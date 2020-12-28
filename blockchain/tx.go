package blockchain

type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

type TxInput struct {
	ID        []byte
	Out       int
	Signature []byte
	PubKey    []byte
}
