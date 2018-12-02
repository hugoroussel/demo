package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]

}

func NewBlock(data string, prevBlockHash []byte) (newBlock *Block) {
	newBlock = &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	newBlock.SetHash()
	fmt.Println("Added new block with data:", data)
	return newBlock
}

func NewGenesisBlock() (firstBlock *Block) {
	return NewBlock("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks", []byte{})
}
