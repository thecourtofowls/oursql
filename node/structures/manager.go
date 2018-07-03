package structures

import (
	"crypto/rand"
	"fmt"

	"github.com/gelembjuk/oursql/lib"
)

// return BlockShort object from bytes
func NewBlockShortFromBytes(bsdata []byte) (*BlockShort, error) {
	bs := &BlockShort{}
	err := bs.DeserializeBlock(bsdata)

	if err != nil {
		return nil, err
	}
	return bs, nil
}

// Make Block object from bytes
func NewBlockFromBytes(bsdata []byte) (*Block, error) {
	bs := &Block{}
	err := bs.DeserializeBlock(bsdata)

	if err != nil {
		return nil, err
	}
	return bs, nil
}

// New "currency" transaction.
func NewTransaction(sqlcommand string, inputs []TXCurrencyInput, outputs []TXCurrrencyOutput) (*Transaction, error) {
	tx := &Transaction{}
	tx.Vin = inputs
	tx.Vout = outputs
	tx.SQLCommand = []byte(sqlcommand)
	return tx, nil
}

// Serialize Transaction
func SerializeTransaction(tx *Transaction) ([]byte, error) {
	// add TX type to know how to deSerialize

	txData, err := tx.serialize()

	if err != nil {
		return nil, err
	}

	return txData, nil
}

// Serialize Transaction
func DeserializeTransaction(txData []byte) (*Transaction, error) {
	// get type from first byte
	tx := &Transaction{}
	err := tx.DeserializeTransaction(txData)

	if err != nil {
		return nil, err
	}
	return tx, nil
}

// New "currency" Coin Base transaction. This transaction must be present in each new block
func NewCoinbaseTransaction(to, data string) (*Transaction, error) {
	if data == "" {
		randData := make([]byte, 20)
		_, err := rand.Read(randData)

		if err != nil {
			return nil, err
		}

		data = fmt.Sprintf("%x", randData)
	}
	tx := &Transaction{}
	txin := TXCurrencyInput{[]byte{}, -1, nil, []byte(data)}
	txout := NewTXOutput(lib.CurrencyPaymentForBlockMade, to)
	tx.Vin = []TXCurrencyInput{txin}
	tx.Vout = []TXCurrrencyOutput{*txout}

	tx.Hash()

	return tx, nil
}

// Sorting of transactions slice
type Transactions []*Transaction

func (c Transactions) Len() int           { return len(c) }
func (c Transactions) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Transactions) Less(i, j int) bool { return c[i].GetTime() < c[j].GetTime() }
