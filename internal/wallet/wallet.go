package wallet

import (
        "crypto/ecdsa"
        "crypto/elliptic"
        "crypto/rand"
        "log"
)

type Wallet struct {
        PrivateKey ecdsa.PrivateKey
        PublicKey  []byte
}

func NewWallet() *Wallet {
        private, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
        if err != nil {
                log.Panic(err)
        }
        pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
        return &Wallet{*private, pubKey}
}