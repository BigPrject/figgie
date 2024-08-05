package main

import (
	"encoding/json"
)

type Message struct {
	Kind string          `json:"kind"`
	Data json.RawMessage `json:"data"`
}

func main() {
	myInv := Inventory{
		spades:   1,
		clubs:    1,
		diamonds: 4,
		hearts:   4,
	}

	myInv.complexPrior()
}
