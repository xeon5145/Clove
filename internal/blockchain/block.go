package blockchain

import (
        "bytes"
        "crypto/sha256"
        "encoding/binary"
        "time"
)

type Block struct {
        Timestamp     int64
        PrevBlockHash []byte
        Hash          []byte
        Data          []byte
        Nonce         int
}

func NewBlock(data []byte, prevBlockHash []byte) *Block {
        block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, data, 0}
        block.SetHash()
        return block
}

func (b *Block) SetHash() {
        headers := bytes.Join(
                [][]byte{
                        b.PrevBlockHash,
                        b.Data,
                        IntToHex(b.Timestamp),
                        IntToHex(int64(b.Nonce)),
                },
                []byte{},
        )
        hash := sha256.Sum256(headers)

        b.Hash = hash[:]
}

func IntToHex(num int64) []byte {
        buff := new(bytes.Buffer)
        err := binary.Write(buff, binary.BigEndian, num)
        if err != nil {
                panic(err)
        }

        return buff.Bytes()
}