package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var tradesMu sync.Mutex

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
			Price:  100,
			Quoter: "",
		},
		Bid: Quote{
			Price:  0,
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
	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		processTrade(update.Trade, gs)
	}()
	go func() {
		defer wg.Done()
		findBBO(update.Spades, gs, "spades")
	}()
	go func() {
		defer wg.Done()
		findBBO(update.Clubs, gs, "clubs")
	}()
	go func() {
		defer wg.Done()
		findBBO(update.Hearts, gs, "hearts")
	}()
	go func() {
		defer wg.Done()
		findBBO(update.Diamonds, gs, "diamonds")
	}()
	wg.Wait()

}

type bboDescription struct {
	book      *Book
	quote     *Quote
	direction string
}

func findBBO(cd CardData, gs *GameState, card string) {
	// should probably make a getter func for the prices...
	// setting the qouter field is probably uncessary for me
	// I should ignore
	var bboDesc bboDescription
	gs.rwMutex.Lock()
	defer gs.rwMutex.Unlock()
	isBBO := false
	switch card {
	case "spades":
		for _, Quote := range cd.Asks {
			if Quote.Price < gs.Orderbook.Spadebook.Ask.Price {
				gs.Orderbook.Spadebook.Ask.Price = Quote.Price
				gs.Orderbook.Spadebook.Ask.Quoter = Quote.Quoter
				isBBO = true
				bboDesc = bboDescription{book: gs.Orderbook.Spadebook, quote: &Quote, direction: "ask"}
			}

		}
		for _, Quote := range cd.Bids {
			if Quote.Price < gs.Orderbook.Spadebook.Ask.Price {
				gs.Orderbook.Spadebook.Bid.Price = Quote.Price
				gs.Orderbook.Spadebook.Bid.Quoter = Quote.Quoter
				isBBO = true
				bboDesc = bboDescription{book: gs.Orderbook.Spadebook, quote: &Quote, direction: "bid"}

			}

		}

	case "clubs":
		for _, Quote := range cd.Asks {
			if Quote.Price < gs.Orderbook.Clubbook.Ask.Price {
				gs.Orderbook.Clubbook.Ask.Price = Quote.Price
				gs.Orderbook.Clubbook.Ask.Quoter = Quote.Quoter
				isBBO = true
				bboDesc = bboDescription{book: gs.Orderbook.Clubbook, quote: &Quote, direction: "ask"}

			}

		}
		for _, Quote := range cd.Bids {
			if Quote.Price > gs.Orderbook.Clubbook.Bid.Price {
				gs.Orderbook.Clubbook.Bid.Price = Quote.Price
				gs.Orderbook.Clubbook.Bid.Quoter = Quote.Quoter
				isBBO = true
				bboDesc = bboDescription{book: gs.Orderbook.Clubbook, quote: &Quote, direction: "bid"}
			}

		}

	case "hearts":
		for _, Quote := range cd.Asks {
			if Quote.Price < gs.Orderbook.Heartbook.Ask.Price {
				gs.Orderbook.Heartbook.Ask.Price = Quote.Price
				gs.Orderbook.Heartbook.Ask.Quoter = Quote.Quoter
				isBBO = true
				bboDesc = bboDescription{book: gs.Orderbook.Heartbook, quote: &Quote, direction: "ask"}
			}

		}
		for _, Quote := range cd.Bids {
			if Quote.Price > gs.Orderbook.Heartbook.Bid.Price {
				gs.Orderbook.Heartbook.Bid.Price = Quote.Price
				gs.Orderbook.Heartbook.Bid.Quoter = Quote.Quoter
				isBBO = true
				bboDesc = bboDescription{book: gs.Orderbook.Heartbook, quote: &Quote, direction: "bid"}
			}

		}

	case "diamonds":
		for _, Quote := range cd.Asks {
			if Quote.Price < gs.Orderbook.Heartbook.Ask.Price {
				gs.Orderbook.Diamondbook.Ask.Price = Quote.Price
				gs.Orderbook.Diamondbook.Ask.Quoter = Quote.Quoter
				isBBO = true
				bboDesc = bboDescription{book: gs.Orderbook.Diamondbook, quote: &Quote, direction: "ask"}

			}

		}
		for _, Quote := range cd.Bids {
			if Quote.Quoter == myplayerName {
				continue
			} else if Quote.Price > gs.Orderbook.Diamondbook.Bid.Price {
				gs.Orderbook.Diamondbook.Bid.Price = Quote.Price
				gs.Orderbook.Diamondbook.Bid.Quoter = Quote.Quoter
				isBBO = true
				bboDesc = bboDescription{book: gs.Orderbook.Diamondbook, quote: &Quote, direction: "bid"}

			}

		}

	}
	// could optimize for bbo ask and buy...
	if isBBO {
		select {
		case bookChannel <- bboDesc:
		default:
		}
	}

}

func clearBook(suit Suit, gs *GameState, price int) {
	//reset every Orderbook execpt suit of trade
	//if time simply by making a map and iterating through and reseting.
	gs.rwMutex.Lock()
	defer gs.rwMutex.Unlock()
	switch suit {
	case spades:
		gs.Orderbook.Clubbook.Ask.Price = 0
		gs.Orderbook.Diamondbook.Ask.Price = 0
		gs.Orderbook.Heartbook.Ask.Price = 0
		gs.Orderbook.Clubbook.Bid.Price = 0
		gs.Orderbook.Diamondbook.Bid.Price = 0
		gs.Orderbook.Heartbook.Bid.Price = 0
		gs.Orderbook.Spadebook.LastPrice = price
	case clubs:
		gs.Orderbook.Spadebook.Ask.Price = 0
		gs.Orderbook.Diamondbook.Ask.Price = 0
		gs.Orderbook.Heartbook.Ask.Price = 0
		gs.Orderbook.Spadebook.Bid.Price = 0
		gs.Orderbook.Diamondbook.Bid.Price = 0
		gs.Orderbook.Heartbook.Bid.Price = 0
		gs.Orderbook.Clubbook.LastPrice = price
	case diamonds:
		gs.Orderbook.Spadebook.Ask.Price = 0
		gs.Orderbook.Clubbook.Ask.Price = 0
		gs.Orderbook.Heartbook.Ask.Price = 0
		gs.Orderbook.Spadebook.Bid.Price = 0
		gs.Orderbook.Clubbook.Bid.Price = 0
		gs.Orderbook.Heartbook.Bid.Price = 0
		gs.Orderbook.Diamondbook.LastPrice = price
	case hearts:
		gs.Orderbook.Spadebook.Ask.Price = 0
		gs.Orderbook.Clubbook.Ask.Price = 0
		gs.Orderbook.Diamondbook.Ask.Price = 0
		gs.Orderbook.Spadebook.Bid.Price = 0
		gs.Orderbook.Clubbook.Bid.Price = 0
		gs.Orderbook.Diamondbook.Bid.Price = 0
		gs.Orderbook.Heartbook.LastPrice = price
	}
}

func processTrade(s string, gs *GameState) {

	var trade Trade
	//check if valid trade
	if s == "" {
		return
	}

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

	clearBook(suit, gs, price)

	trade = Trade{
		Card:   suit,
		Price:  price,
		Buyer:  arr[2],
		Seller: arr[3],
	}
	if trade.Buyer == myplayerName {
		gs.Inventory.update(trade.Card, true)
	} else if trade.Seller == myplayerName {
		gs.Inventory.update(trade.Card, false)
	}
	appendTrade(&trade, gs)

}

// optimize here

func appendTrade(newTrade *Trade, gs *GameState) {
	gs.rwMutex.Lock()
	gs.Trades = append(gs.Trades, *newTrade)
	gs.rwMutex.Unlock()
	select {
	case tradeChannel <- struct{}{}:
	default:
	}
}
