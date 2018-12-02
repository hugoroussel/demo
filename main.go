package main

import (
	"fmt"
	"strconv"
)

func main() {
	Blockchain := NewBlockchain()

	Blockchain.AddBlock("Send 1 coin to Patrick")
	Blockchain.AddBlock("Send 2 coins to Martin")

	for _, block := range Blockchain.blocks {
		block.Display()
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
