package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var tradesMu sync.Mutex

var tradeChannel = make(chan struct{}, 1)

type Quote struct {
	Price  int    `json:"price"`
	Quoter string `json:"player_name"`
}

type Book struct {
	Ask       Quote
	Bid       Quote
	LastPrice int
}

type Orderbook struct {
	Spadebook   *Book
	Clubbook    *Book
	Heartbook   *Book
	Diamondbook *Book
}

func (book *Book) UpdateLastPrice(p int) {
	book.LastPrice = p
}

func NewBook() *Book {
	return &Book{
		Ask: Quote{
			Price:  0,
			Quoter: "",
		},
		Bid: Quote{
			Price:  99,
			Quoter: "",
		},
		LastPrice: 0,
	}
}

func NewOrderbook() *Orderbook {
	return &Orderbook{
		Spadebook:   NewBook(),
		Clubbook:    NewBook(),
		Heartbook:   NewBook(),
		Diamondbook: NewBook(),
	}
}

func processUpdate(update UpdateStruct, gs *GameState) {
	processTrade(update.Trade, gs)
	// finish out everything else
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
	appendTrade(&trade, gs)
}

// optimize here

func appendTrade(newTrade *Trade, gs *GameState) {
	gs.mutex.Lock()
	gs.Trades = append(gs.Trades, *newTrade)
	gs.mutex.Unlock()
	select {
	case tradeChannel <- struct{}{}:
	default:
	}
}
