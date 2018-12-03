package main

import (
	"fmt"
	"strconv"
)

func main() {
	Blockchain := NewBlockchain()

	Blockchain.AddBlock("Send 1 coin to Patrick")
	Blockchain.AddBlock("Send 2 coins to Martin")
	Blockchain.AddBlock("Send 5 coins to George")

	for _, block := range Blockchain.blocks {
		if block.Number == 1 {
			block.Hash = []byte("hello!")

		}
		block.Display()
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
