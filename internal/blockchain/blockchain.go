package blockchain

type Blockchain struct {
        Blocks []*Block
}

func NewGenesisBlock() *Block {
        return NewBlock([]byte("Genesis Block"), []byte{})
}

func NewBlockchain() *Blockchain {
        return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func (bc *Blockchain) AddBlock(data []byte) {
        prevBlock := bc.Blocks[len(bc.Blocks)-1]
        newBlock := NewBlock(data, prevBlock.Hash)
        bc.Blocks = append(bc.Blocks, newBlock)
}