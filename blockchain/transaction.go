package blockchain

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}
