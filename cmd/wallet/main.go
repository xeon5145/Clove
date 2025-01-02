package main

import (
        "fmt"
        "clove/internal/wallet"
)

func main() {
        w := wallet.NewWallet()
        fmt.Printf("Private Key: %x\n", w.PrivateKey.D.Bytes())
        fmt.Printf("Public Key: %x\n", w.PublicKey)
}