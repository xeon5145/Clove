package main

import (
	"fmt"
	"blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Third Block")

	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d, Hash: %s, PrevHash: %s, Data: %s\n", block.Index, block.Hash, block.PrevHash, block.Data)
	}
}
