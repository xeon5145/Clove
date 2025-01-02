package blockchain

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int
	Timestamp time.Time
	Data      string
	PrevHash  string
	Hash      string
}

func (b *Block) DeriveHash() {
	info := string(b.Index) + b.Timestamp.String() + b.Data + b.PrevHash
	hash := sha256.Sum256([]byte(info))
	b.Hash = hex.EncodeToString(hash[:])
}
