package main

import "encoding/json"

type Suit string

type Trade struct {
	Card   Suit
	Price  int
	Buyer  string
	Seller string
}

type Message struct {
	Kind string          `json:"kind"`
	Data json.RawMessage `json:"data"`
}

var players map[string]int = map[string]int{"me": 0, "p2": 1, "p3": 2, "p4": 3}
var Trades = make([]Trade, 0, 100)

const (
	spades     Suit   = "spades"
	clubs      Suit   = "clubs"
	diamonds   Suit   = "diamonds"
	hearts     Suit   = "hearts"
	dealing    string = "dealing_cards"
	update     string = "update"
	endOfRound string = "end_round"
	endOfGame  string = "end_game"
)

const (
	wsURL      = "ws://testnet-ws.figgiewars.com"
	resURL     = "http://testnet.figgiewars.com"
	playerID   = "LebronJames"
	playerName = "bellamy"
)

// Helpers

func (s Suit) getGoalSuit() Suit {
	switch s {
	case spades:
		return clubs
	case clubs:
		return spades
	case diamonds:
		return hearts
	case hearts:
		return diamonds
	}

	return ""
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
