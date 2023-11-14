package main

func main() {
	// log.Println("test")
	// b := NewBlock(1, "fdad")
	// //fmt.Println(b)
	// b.Print(os.Stdout)

	// blockChain := NewBlockchain()
	// blockChain.Print()
	// blockChain.CreateBlock(5, "hash 1")
	// blockChain.Print()
	// blockChain.CreateBlock(2, "hash 2")
	// blockChain.Print()
	// blockChain.CreateBlock(2, "hash 3")
	// blockChain.Print()

	// blockChain := NewBlockchain()
	// blockChain.Print()

	// previousHash := blockChain.LastBlock().Hash()
	// blockChain.CreateBlock(5, previousHash)
	// blockChain.Print()

	// previousHash = blockChain.LastBlock().Hash()
	// blockChain.CreateBlock(2, previousHash)
	// blockChain.Print()
	blockChain := NewBlockchain()
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 1.0)
	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, previousHash)
	blockChain.Print()

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("X", "Y", 3.0)
	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, previousHash)
	blockChain.Print()
}
