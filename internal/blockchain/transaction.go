package blockchain

type Transaction struct {
        ID      []byte
        Inputs  []TXInput
        Outputs []TXOutput
}

type TXInput struct {
        Txid      []byte
        Vout      int
        ScriptSig string
}

type TXOutput struct {
        Value         int
        ScriptPubKey string
}