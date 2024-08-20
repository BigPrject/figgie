package main

import "sync"

type GameState struct {
	Inventory     *Inventory
	Orderbook     *Orderbook
	Trades        []Trade
	Probabilities map[Suit]float64
	mutex         sync.RWMutex
	Balance       int
	*Fundbot
}

func NewGameState() *GameState {
	return &GameState{
		Inventory:     &Inventory{},
		Orderbook:     newBook(),
		Trades:        make([]Trade, 100),
		Probabilities: make(map[Suit]float64),
		Balance:       0,
	}
}

/*
 func listenToUpdates(gs *GameState, fd *Fundbot) {
	for {
		select {
		case <-updateChannel:
			fd.runFundamental(gs)
		}
	}
}

*/
