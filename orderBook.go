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

func newBook() Orderbook {

	return Orderbook{
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

func processUpdate(update UpdateStruct) {
	processTrade(update.Trade)

}

func processTrade(s string) {

	var trade Trade

	arr := strings.Split(s, ",")

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
	appendTrade(trade)
}

func appendTrade(newTrade Trade) {
	tradesMu.Lock()
	Trades = append(Trades, newTrade)
	tradesMu.Unlock()
	// Signal that an update has occurred
	select {
	case updateChannel <- struct{}{}:
	default:
	}
}
