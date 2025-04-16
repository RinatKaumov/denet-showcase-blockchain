package main

import (
	"denetshow/core"
	"fmt"
)

func main() {
	chain := core.NewBlockchain()

	chain.AddBlock([]core.Transaction{
		core.NewTransaction("Alice", "Bob", 10),
		core.NewTransaction("Bob", "Charlie", 5),
	})

	chain.AddBlock([]core.Transaction{
		core.NewTransaction("Charlie", "Alice", 2),
	})

	for _, block := range chain.Blocks {
		fmt.Println("---------------")
		fmt.Printf("Height: %d\n", block.Header.Height)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("PrevHash: %x\n", block.Header.PrevHash)
		fmt.Printf("Nonce: %d\n", block.Header.Nonce)
	}
}
