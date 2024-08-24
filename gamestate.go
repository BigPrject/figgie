package main

import "sync"

type GameState struct {
	Inventory *Inventory
	Orderbook *Orderbook
	Trades    []Trade
	rwMutex   sync.RWMutex
	Balance   int
}

func NewGameState() *GameState {
	return &GameState{
		Inventory: &Inventory{},
		Orderbook: NewOrderbook(),
		Trades:    make([]Trade, 100),
		Balance:   350,
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
