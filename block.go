package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"
)

func init() {
	log.SetPrefix("block: ")
}

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

func NewBlock(nonce int, prevHash [32]byte, txs []*Transaction) *Block {
	return &Block{
		nonce:        nonce,
		timestamp:    time.Now().UnixNano(),
		previousHash: prevHash,
		transactions: txs,
	}
}

func (b *Block) String() string {
	if b == nil {
		return ""
	}
	return fmt.Sprintf("%-15v: [%v]\n%-15v: [%v]\n%-15v: [%x]\n%-15v: %v\n",
		"timestamp", b.timestamp, "nonce", b.nonce, "prevHash", b.previousHash, "transaction", b.transactions)
}

func (b *Block) Print(w io.Writer) {
	fmt.Fprint(w, b)
	// for _, t := range b.transactions {
	// 	t.Print(w)
	// }
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transaction"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}
