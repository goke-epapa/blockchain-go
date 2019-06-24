package main

import "github.com/boltdb/bolt"

const (
	dbFile              = "blockchain_db"
	lastBlockIdentifier = "l"
	// BlocksBucket block DB identifier
	BlocksBucket = "blocks"
)

// Blockchain blockchain struct
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

// AddBlock adds a new block to the chain
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlocksBucket))
		lastHash = b.Get([]byte(lastBlockIdentifier))

		return nil
	})

	newBlock := NewBlock(data, lastHash)

	err := bd.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialise())
		err = b.Put([]byte(lastBlockIdentifier), newBlock.Hash)
		bc.tip = newBlock.Hash

		return nil
	})
}

// NewBlockchain creates a new blockchain
func NewBlockchain() *Blockchain {
	var tip []byte
	// TODO handle errors
	db, _ := bolt.Open(dbFile, 0600, nil)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, _ := tx.CreateBucket([]byte(BlocksBucket))

			sg, _ := genesis.Serialise()
			_ = b.Put(genesis.Hash, sg)
			_ = b.Put([]byte(lastBlockIdentifier), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte(lastBlockIdentifier))
		}

		return nil
	})

	bc := &Blockchain{tip, db}

	return bc
}

// Iterator return a new blockchain iterator
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}
