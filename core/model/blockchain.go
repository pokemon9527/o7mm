package model


type BlockChain struct {
  blocks []*Block
}


func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// 创世块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
