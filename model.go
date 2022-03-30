package main

type Transaction struct {
	TxId string             `json:"txid"`
	VIn  []InputTransaction `json:"vin"`
}

type InputTransaction struct {
	TxId string `json:"txid"`
}

type kv struct {
	Key   string
	Value int
}
