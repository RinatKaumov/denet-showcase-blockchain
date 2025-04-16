package core

import (
	"bytes"
	"crypto/sha256"
	"denetshow/crypto"
	"encoding/gob"
	"time"
)

type Header struct {
	Height    uint32
	Timestamp int64
	PrevHash  []byte
	DataHash  []byte
	Nonce     uint64
}

type Block struct {
	Header       *Header
	Transactions []Transaction
	Hash         []byte
}

func NewBlock(transactions []Transaction, prevHash []byte, height uint32) *Block {
	timestamp := time.Now().Unix()
	dataHash := hashTransactions(transactions)

	header := &Header{
		Height:    height,
		Timestamp: timestamp,
		PrevHash:  prevHash,
		DataHash:  dataHash,
		Nonce:     0,
	}

	block := &Block{
		Header:       header,
		Transactions: transactions,
		Hash:         nil,
	}

	// ✅ Используем новую структуру вместо []byte
	powInput := crypto.BlockHeaderData{
		PrevHash:  prevHash,
		DataHash:  dataHash,
		Timestamp: timestamp,
		Height:    height,
	}

	pow := crypto.NewProofOfWork(powInput)
	nonce, hash := pow.RunParallel()

	header.Nonce = nonce
	block.Hash = hash

	return block
}

func hashTransactions(txs []Transaction) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	_ = enc.Encode(txs)
	hash := sha256.Sum256(buf.Bytes())
	return hash[:]
}
