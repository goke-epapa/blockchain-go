package main

import (
	"github.com/boltdb/bolt"
)

// BlockchainIterator blockchain iterator struct
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// Next gets the next block
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserialiseBlock(encodedBlock)

		return nil
	})

	i.currentHash = block.PrevBlockHash

	return block
}

// Iterable interface
type Iterable interface {
	Iterator() BlockchainIterator
}
