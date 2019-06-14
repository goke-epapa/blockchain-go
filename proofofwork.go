package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 24
const maxNonce = math.MaxInt64

// ProofOfWork a struct representing proof of work
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork creates a new ProofOfWork
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			convertIntToHex(pow.block.Timestamp),
			convertIntToHex(int64(targetBits)),
			convertIntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func convertIntToHex(number int64) []byte {
	return []byte(fmt.Sprintf("%x", number))
}

// Run runs the ProofOfWork
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println("")

	return nonce, hash[:]
}
