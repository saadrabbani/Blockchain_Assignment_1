package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Block struct {
	Index       int
	Timestamp   string
	transaction string
	Hash        string
	PrevHash    string
	nonce       int
}

func (b *Block) CreateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.transaction + b.PrevHash + strconv.Itoa(b.nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func NewBlock(transaction string, nonce int, previousHash string) Block {
	t := time.Now()
	newBlock := Block{rand.Intn(100), t.String(), transaction, "", previousHash, nonce}
	newBlock.Hash = newBlock.CreateHash()
	return newBlock
}

func ListBlocks(blocks []Block) {
	for _, block := range blocks {
		fmt.Printf("Index: %d Timestamp: %s Transaction: %s Hash: %s PrevHash: %s Nonce: %d \n", block.Index, block.Timestamp, block.transaction, block.Hash, block.PrevHash, block.nonce)
		fmt.Println()
	}
}

func ChangeBlock(block *Block, transaction string) {
	block.transaction = transaction
	block.Hash = ""
	block.Hash = block.CreateHash()
	fmt.Println("The new details of the block are as follows:")
	fmt.Printf("Index: %d Timestamp: %s Transaction: %s Hash: %s PrevHash: %s Nonce: %d \n", block.Index, block.Timestamp, block.transaction, block.Hash, block.PrevHash, block.nonce)
	fmt.Println("The Block has been changed")
}

func VerifyChain(Blockchain []Block) {

	for i := 1; i < len(Blockchain); i++ {
		if Blockchain[i].PrevHash != Blockchain[i-1].Hash {
			fmt.Println("Blockchain is not valid, block %d is not valid", i-1)
			fmt.Println("The details of the block are as follows:")
			fmt.Printf("Index: %d Timestamp: %s Transaction: %s Hash: %s PrevHash: %s Nonce: %d \n", Blockchain[i-1].Index, Blockchain[i-1].Timestamp, Blockchain[i-1].transaction, Blockchain[i-1].Hash, Blockchain[i-1].PrevHash, Blockchain[i-1].nonce)
			fmt.Println()
			return
		}
	}
	fmt.Println("Blockchain is valid")
}

func main() {
	t := time.Now()
	genesisBlock := Block{0, t.String(), "Genesis Block", "", "", 0}
	genesisBlock.Hash = genesisBlock.CreateHash()
	Blockchain := []Block{genesisBlock}

	max := 9999
	min := 1000

	for i := 1; i < 10; i++ {
		newBlock := NewBlock("Random Transaction", rand.Intn(max-min)+min, Blockchain[i-1].Hash)
		Blockchain = append(Blockchain, newBlock)
	}
	ListBlocks(Blockchain)
	VerifyChain(Blockchain)
	ChangeBlock(&Blockchain[3], "Changed Block")

	VerifyChain(Blockchain)
	ListBlocks(Blockchain)
}
