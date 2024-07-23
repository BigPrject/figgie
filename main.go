package main

import (
	"encoding/json"
)

type Message struct {
	Kind string          `json:"kind"`
	Data json.RawMessage `json:"data"`
}

func main() {

}
