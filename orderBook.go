package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var tradesMu sync.Mutex

var updateChannel = make(chan struct{}, 1)

type Qoute struct {
	Price  int    `json:"price"`
	Qouter string `json:"player_name"`
}

type Orderbook struct {
	ask Qoute
	bid Qoute

	lastPrice int
}

func (book *Orderbook) newlastPrice(p int) {
	book.lastPrice = p
}

func newBook() *Orderbook {

	return &Orderbook{
		ask: Qoute{
			Price:  0,
			Qouter: "",
		},
		bid: Qoute{
			Price:  99,
			Qouter: "",
		},
		lastPrice: 0,
	}
}

func processUpdate(update UpdateStruct, gs *GameState) {
	processTrade(update.Trade, gs)

}

func processTrade(s string, gs *GameState) {

	var trade Trade

	arr := strings.Split(s, ",")
	if len(arr) != 4 {
		fmt.Println("Invalid trade format:", s)
		return
	}

	price, err := strconv.Atoi(arr[1])
	if err != nil {
		fmt.Println("Error converting price:", err)
	}

	var suit Suit
	switch arr[0] {
	case "spades":
		suit = spades
	case "clubs":
		suit = clubs
	case "diamonds":
		suit = diamonds
	case "hearts":
		suit = hearts
	default:
		fmt.Println("Invalid card suit")
	}

	trade = Trade{
		Card:   suit,
		Price:  price,
		Buyer:  arr[2],
		Seller: arr[3],
	}
	appendTrade(trade, gs)
}

func appendTrade(newTrade Trade, gs *GameState) {
	gs.mutex.Lock()
	gs.Trades = append(gs.Trades, newTrade)
	gs.mutex.Unlock()
	select {
	case updateChannel <- struct{}{}:
	default:
	}
}
