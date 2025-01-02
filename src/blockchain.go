package blockchain
import "time"

type Blockchain struct {
        Blocks []Block
}

func (bc *Blockchain) NewBlockchain() *Blockchain {
        genesisBlock := Block{}
        genesisBlock.Index = 0
        genesisBlock.Timestamp = time.Now()
        genesisBlock.Data = "Genesis Block"
        genesisBlock.PrevHash = "" 
        genesisBlock.DeriveHash()

        bc.Blocks = append(bc.Blocks, genesisBlock)
        return bc
}

func (bc *Blockchain) AddBlock(data string) {
        prevBlock := bc.Blocks[len(bc.Blocks)-1] 
        newBlock := Block{ 
                Index:     len(bc.Blocks) + 1, 
                Timestamp: time.Now(), 
                Data:      data, 
                PrevHash:  prevBlock.Hash, 
        }
        newBlock.DeriveHash() 
        bc.Blocks = append(bc.Blocks, newBlock) 
}