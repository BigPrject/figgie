package main

import "encoding/json"

// Channels
var tradeChannel = make(chan struct{}, 1)
var bookChannel = make(chan bboDescription, 1)
var invChannel = make(chan struct{}, 1)
var bboChannel = make(chan struct{}, 1)

type Suit string

type Order struct {
	Card      string `json:"card"`
	Price     int    `json:"price"`
	Direction string `json:"direction"`
}
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

// add seprate map of players on exchange, where each of them correspond p(i)
var players map[string]int = map[string]int{"me": 0, "p2": 1, "p3": 2, "p4": 3}

//var Trades = make([]Trade, 0, 100)

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
	wsURL        = "ws://testnet-ws.figgiewars.com"
	resURL       = "http://testnet.figgiewars.com"
	playerID     = "LebronJames"
	myplayerName = "bellamy"
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
