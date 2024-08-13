package main

import "sync"

type GameState struct {
	Inventory     *Inventory
	Orderbook     *Orderbook
	Trades        []Trade
	Probabilities map[Suit]float64
	mutex         sync.RWMutex
	Balance       int
}

func NewGameState() *GameState {
	return &GameState{
		Inventory:     &Inventory{},
		Orderbook:     newBook(),
		Trades:        make([]Trade, 0),
		Probabilities: make(map[Suit]float64),
		Balance:       0,
	}
}
