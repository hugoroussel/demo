package main

import (
	"fmt"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func NewBlock(data string, prevBlockHash []byte) (newBlock *Block) {
	newBlock = &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(newBlock)
	nonce, hash := pow.Run()
	newBlock.Nonce = nonce
	newBlock.Hash = hash
	//newBlock.SetHash()
	//fmt.Println("Added new block with data:", data)
	return newBlock
}

func NewGenesisBlock() (firstBlock *Block) {
	return NewBlock("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks", []byte{})
}

func (b *Block) Display() {
	fmt.Println("Timestamp: ", b.Timestamp)
	fmt.Printf("Data: %s \n", string(b.Data))
	fmt.Printf("Prev. Block Hash: %x \n", b.PrevBlockHash)
	fmt.Printf("Block Hash: %x \n", b.Hash)
	fmt.Println()
}

/*
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]

}*/
