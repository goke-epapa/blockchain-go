package main

import "github.com/boltdb/bolt"

const (
	dbFile       = "blockchain_db"
	blocksBucket = "blocks"
)

// Blockchain blockchain struct
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

// AddBlock adds a new block to the chain
func (bc *Blockchain) AddBlock(data string) {
	// prevBlock := bc.blocks[len(bc.blocks)-1]
	// newBlock := NewBlock(data, prevBlock.Hash)
	// bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates a new blockchain
func NewBlockchain() *Blockchain {
	var tip []byte
	// TODO handle errors
	db, _ := bolt.Open(dbFile, 0600, nil)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, _ := tx.CreateBucket([]byte(blocksBucket))

			sg, _ := genesis.Serialise()
			_ = b.Put(genesis.Hash, sg)
			_ = b.Put([]byte("g"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("g"))
		}

		return nil
	})

	bc := &Blockchain{tip, db}

	return bc
}
