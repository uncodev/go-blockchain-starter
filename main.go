package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Index     int
	Timestamp int64
	Data      string
	PrevHash  string
	Hash      string
}

// Blockchain represents the blockchain itself
type Blockchain struct {
	Chain []Block
}

// CalculateHash calculates the hash of a block
func calculateHash(index int, timestamp int64, data string, prevHash string) string {
	payload := fmt.Sprintf("%d%d%s%s", index, timestamp, data, prevHash)
	hash := sha256.Sum256([]byte(payload))
	return hex.EncodeToString(hash[:])
}

// CreateGenesisBlock creates the first block (genesis block) in the blockchain
func createGenesisBlock() Block {
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().Unix(),
		Data:      "Genesis Block",
		PrevHash:  "",
	}
	genesisBlock.Hash = calculateHash(genesisBlock.Index, genesisBlock.Timestamp, genesisBlock.Data, genesisBlock.PrevHash)
	return genesisBlock
}

// NewBlockchain creates a new blockchain with the genesis block
func NewBlockchain() *Blockchain {
	genesisBlock := createGenesisBlock()
	return &Blockchain{Chain: []Block{genesisBlock}}
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := mineBlock(prevBlock, data)
	bc.Chain = append(bc.Chain, newBlock)
}

// MineBlock mines a new block using proof-of-work
func mineBlock(prevBlock Block, data string) Block {
	index := prevBlock.Index + 1
	timestamp := time.Now().Unix()
	prevHash := prevBlock.Hash
	hash := calculateHash(index, timestamp, data, prevHash)
	return Block{
		Index:     index,
		Timestamp: timestamp,
		Data:      data,
		PrevHash:  prevHash,
		Hash:      hash,
	}
}

func main() {
	bc := NewBlockchain()

	// Add blocks to the blockchain
	bc.AddBlock("Transaction 1")
	bc.AddBlock("Transaction 2")

	// Print the blockchain
	for _, block := range bc.Chain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("---------------------")
	}
}
