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
		spades:   4,
		clubs:    4,
		diamonds: 1,
		hearts:   1,
	}

	myInv.calcPrior()
}
