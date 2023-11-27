package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

// Transaction represents a transaction in a block.
type Transaction struct {
	Data []byte
}

func TransactionsToString(trans []*Transaction) string {
	var transactionStrings []string

	for _, t := range trans {
		transactionStrings = append(transactionStrings, string(t.Data))
	}

	return fmt.Sprintf("[%s]", strings.Join(transactionStrings, ","))
}

// Block represents a block in the blockchain.
type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) DeriveHash() {
	var transactionsData []byte
	for _, tx := range b.Transactions {
		transactionsData = append(transactionsData, []byte(tx.Data)...)
	}
	info := bytes.Join(
		[][]byte{
			[]byte(time.Unix(b.Timestamp, 0).Format(time.RFC3339)),
			transactionsData,
			b.PrevBlockHash,
		},
		[]byte{},
	)
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock creates a new block.
func CreateBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
	}
	block.DeriveHash()
	return block
}

// Genesis creates the genesis block.
func Genesis() *Block {
	return CreateBlock([]*Transaction{
		{Data: []byte("Genesis")},
	}, []byte{})
}

// Blockchain represents the blockchain.
type Blockchain struct {
	blocks []*Block
}

// AddBlock adds a new block to the blockchain.
func (chain *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(transactions, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

// Init.
func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}
