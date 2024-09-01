package main

import "sync"

type GameState struct {
	Inventory *Inventory
	Orderbook *Orderbook
	Trades    []Trade
	myTrades  []Trade
	rwMutex   sync.RWMutex
	Balance   int
	players   map[string]int
	goalSuit  Suit
	arbPrice  int
}

func NewGameState() *GameState {
	return &GameState{
		Inventory: &Inventory{},
		Orderbook: NewOrderbook(),
		Trades:    make([]Trade, 100),
		Balance:   350,
		players:   make(map[string]int, 4),
		myTrades:  make([]Trade, 10),
		arbPrice:  0,
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

func getHand(gs *GameState) *map[Suit]int {
	return &map[Suit]int{spades: gs.Inventory.Spades,
		clubs:    gs.Inventory.Clubs,
		diamonds: gs.Inventory.Diamonds, hearts: gs.Inventory.Hearts}
}
