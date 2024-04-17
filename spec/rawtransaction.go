package spec

type BtcRawTransaction struct {
	Txid          string `json:"txid"`
	Hash          string `json:"hash"`
	Version       int32  `json:"version"`
	Size          int    `json:"size"`
	Vsize         int    `json:"vsize"`
	Weight        int    `json:"weight"`
	Locktime      int    `json:"locktime"`
	Vin           []Vin  `json:"vin"`
	Vout          []Vout `json:"vout"`
	Hex           string `json:"hex"`
	Blockhash     string `json:"blockhash"`
	Confirmations uint64 `json:"confirmations"`
	Time          int64  `json:"time"`
	Blocktime     int64  `json:"blocktime"`
}

type Vin struct {
	Txid        string    `json:"txid"`
	Vout        int       `json:"vout"`
	ScriptSig   ScriptSig `json:"scriptSig"`
	Txinwitness []string  `json:"txinwitness"`
	PrevOut     PrevOut   `json:"prevOut"`
	Sequence    uint32    `json:"sequence"`
}

type ScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

type PrevOut struct {
	Generated    bool               `json:"generated"`
	Height       int                `json:"height"`
	Value        float64            `json:"value"`
	ScriptPubKey ScriptPubKeyResult `json:"scriptPubKey"`
}

type ScriptPubKeyResult struct {
	Asm     string `json:"asm"`
	Desc    string `json:"desc"`
	Hex     string `json:"hex"`
	Address string `json:"address"`
	Type    string `json:"type"`
}

type Vout struct {
	Value        float64            `json:"value"`
	N            int                `json:"n"`
	ScriptPubKey ScriptPubKeyResult `json:"scriptPubKey"`
}
