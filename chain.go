package main

import (
	"fmt"
	"os"
	"strings"
)

type Chain struct {
	txPool []*Transaction
	chain  []*Block
}

func (c *Chain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, c.txPool)
	c.chain = append(c.chain, b)
	c.txPool = []*Transaction{}
	return b
}

func NewBlockchain() *Chain {
	b := &Block{}
	bc := new(Chain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Chain) Print() {
	fmt.Printf("%s  Begin  %s", strings.Repeat("=", 25), strings.Repeat("=", 25))
	for i, block := range bc.chain {
		fmt.Printf("\n%s Chain %d %s\n", strings.Repeat("=", 25), i,
			strings.Repeat("=", 25))
		block.Print(os.Stdout)
	}
	fmt.Printf("%s  END   %s\n\n\n", strings.Repeat("=", 25), strings.Repeat("=", 25))
}

func (c *Chain) LastBlock() *Block {
	return c.chain[len(c.chain)-1]
}

func (c *Chain) AddTransaction(sender, recipient string, value float32) {
	c.txPool = append(c.txPool, NewTransaction(sender, recipient, value))
}
