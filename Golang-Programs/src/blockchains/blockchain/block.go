package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// a chain of blocks
type BlockChain struct {
	Blocks []*Block
}

// each block gonna have a data along with a hash
// this block will attatch to last (previous) block
type Block struct {
	Hash []byte    // hash of this block
	Data []byte    // data (ledger, doc ets) in this block 
	PrevHash []byte // last block's hash 
}

// to get hash value
func (b *Block) DeriveHash() {
	// join bytes' slices
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info) // actual blockchain hash is far complex
	b.Hash = hash[:]

	// here hashes are derived using this data + "previous hash"
	// any change in previous data will change its hash and 
	// it'll modify all following blocks' hash too
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
