package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

// a chain of blocks
type BlockChain struct {
	LastHash []byte
	Database []*badger.DB
}

// to iterate through blocks to view
type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

func InitBlockChain() *BlockChain {
	var lastHash []byte

	opts := badger.DefaultOptions
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	err := db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badgerErr.ErrKeyNotFound { // "lh" key is for last-hash
			fmt.Println("No existing blockchain found")
			genesis := Genesis()
			fmt.Println("Genesis proved")
			err = txn.Set(genesis.Hash, genesis.Serialize()) // putting in the DB
			Handle(err)
			err = txn.Set([]bytes("lh"), genesis.Hash)

			lastHash = genesis.Hash

			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			Handle(err)
			lastHash, err = item.Value()
			return err
		}
	})

	Handle(err)

	blockchain := BlockChain{lastHash, db}
	return &blockchain
}

func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get("lh")
		Handle(err)

		lastHash, err = item.Value()

		return err
	})

	Handle(err)

	newBlock := CreateBlock(data, lastHash)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.hash, newBlock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err
	})
	Handle(err)
}

// inplementing an iterator for BlockChain on badger DB
func (chain *BlockChain) Iterator() *BlockChainIterator {
	iter := &BlockChainIterator{chain.LastHash, chain.Database}

	return iter
}

func (iter *BlockChainIterator) Next() *Block {
	var block *Block

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		Handle(err)

		encodedBlock, err = item.Value()
		block = Deserialize(encodedBlock)

		return err
	})
	Handle(err)

	iter.CurrentHash = block.PrevHash

	return block
}
