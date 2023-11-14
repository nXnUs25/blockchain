package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Transaction struct {
	sender_blockchain_address    string
	recipient_blockchain_address string
	value                        float32
}

func NewTransaction(sender, recipient string, value float32) *Transaction {
	return &Transaction{
		sender_blockchain_address:    sender,
		recipient_blockchain_address: recipient,
		value:                        value,
	}
}

func (t *Transaction) String() string {
	if t == nil {
		return ""
	}
	return fmt.Sprintf("\n%v\n%-35v: [%v]\n%-35v: [%v]\n%-35v: [%.1f]\n%[1]v\n", strings.Repeat("-", 50),
		"SenderBlockChainAddress", t.sender_blockchain_address, "RecipientBlockChainAddress", t.recipient_blockchain_address, "Value", t.value)
}

func (t *Transaction) Print(w io.Writer) {
	fmt.Fprintln(w, strings.Repeat("-", 40))
	fmt.Fprint(w, t)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.sender_blockchain_address,
		Recipient: t.recipient_blockchain_address,
		Value:     t.value,
	})
}
