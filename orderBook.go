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

	go processTrade(update.Trade, gs)

	go makeBook(update, gs)
}

func makeBook(update UpdateStruct, gs *GameState) {
	go findBBO(update.Spades, gs, "spades")
	go findBBO(update.Clubs, gs, "clubs")
	go findBBO(update.Hearts, gs, "hearts")
	go findBBO(update.Diamonds, gs, "diamonds")
	// add waitgroup
}

func findBBO(cd CardData, gs *GameState, card string) {
	// should probably make a getter func for the prices...
	// setting the qouter field is probably uncessary for me
	switch card {
	case "spades":
		for _, Quote := range cd.Asks {
			if Quote.Price < gs.Orderbook.Spadebook.Ask.Price {
				gs.Orderbook.Spadebook.Ask.Price = Quote.Price
				gs.Orderbook.Spadebook.Ask.Quoter = Quote.Quoter
			}

		}
		for _, Quote := range cd.Bids {
			if Quote.Price < gs.Orderbook.Spadebook.Ask.Price {
				gs.Orderbook.Spadebook.Bid.Price = Quote.Price
				gs.Orderbook.Spadebook.Bid.Quoter = Quote.Quoter
			}

		}

	case "clubs":
		for _, Quote := range cd.Asks {
			if Quote.Price < gs.Orderbook.Clubbook.Ask.Price {
				gs.Orderbook.Clubbook.Ask.Price = Quote.Price
				gs.Orderbook.Clubbook.Ask.Quoter = Quote.Quoter

			}

		}
		for _, Quote := range cd.Bids {
			if Quote.Price > gs.Orderbook.Clubbook.Bid.Price {
				gs.Orderbook.Clubbook.Bid.Price = Quote.Price
				gs.Orderbook.Clubbook.Bid.Quoter = Quote.Quoter

			}

		}

	case "hearts":
		for _, Quote := range cd.Asks {
			if Quote.Price < gs.Orderbook.Heartbook.Ask.Price {
				gs.Orderbook.Heartbook.Ask.Price = Quote.Price
				gs.Orderbook.Heartbook.Ask.Quoter = Quote.Quoter

			}

		}
		for _, Quote := range cd.Bids {
			if Quote.Price > gs.Orderbook.Heartbook.Bid.Price {
				gs.Orderbook.Heartbook.Bid.Price = Quote.Price
				gs.Orderbook.Heartbook.Bid.Quoter = Quote.Quoter

			}

		}

	case "diamonds":
		for _, Quote := range cd.Asks {
			if Quote.Price < gs.Orderbook.Heartbook.Ask.Price {
				gs.Orderbook.Diamondbook.Ask.Price = Quote.Price
				gs.Orderbook.Diamondbook.Ask.Quoter = Quote.Quoter

			}

		}
		for _, Quote := range cd.Bids {
			if Quote.Price > gs.Orderbook.Diamondbook.Bid.Price {
				gs.Orderbook.Diamondbook.Bid.Price = Quote.Price
				gs.Orderbook.Diamondbook.Bid.Quoter = Quote.Quoter
			}

		}

	}

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
