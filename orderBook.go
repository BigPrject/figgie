package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

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
	suit      Suit
	quote     *Quote
	direction string
}

func findBBO(cd CardData, gs *GameState, card string) {
	// should probably make a getter func for the prices...
	// setting the qouter field is probably uncessary for me
	// I should ignore
	//does bbo even matter to me?
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
				bboDesc = bboDescription{suit: spades, quote: &Quote, direction: "ask"}
			}

		}
		for _, Quote := range cd.Bids {
			if Quote.Price < gs.Orderbook.Spadebook.Ask.Price {
				gs.Orderbook.Spadebook.Bid.Price = Quote.Price
				gs.Orderbook.Spadebook.Bid.Quoter = Quote.Quoter
				isBBO = true
				bboDesc = bboDescription{suit: spades, quote: &Quote, direction: "bid"}

			}

		}

	case "clubs":
		for _, Quote := range cd.Asks {
			if Quote.Price < gs.Orderbook.Clubbook.Ask.Price {
				gs.Orderbook.Clubbook.Ask.Price = Quote.Price
				gs.Orderbook.Clubbook.Ask.Quoter = Quote.Quoter
				isBBO = true
				bboDesc = bboDescription{suit: clubs, quote: &Quote, direction: "ask"}

			}

		}
		for _, Quote := range cd.Bids {
			if Quote.Price > gs.Orderbook.Clubbook.Bid.Price {
				gs.Orderbook.Clubbook.Bid.Price = Quote.Price
				gs.Orderbook.Clubbook.Bid.Quoter = Quote.Quoter
				isBBO = true
				bboDesc = bboDescription{suit: clubs, quote: &Quote, direction: "bid"}
			}

		}

	case "hearts":
		for _, Quote := range cd.Asks {
			if Quote.Price < gs.Orderbook.Heartbook.Ask.Price {
				gs.Orderbook.Heartbook.Ask.Price = Quote.Price
				gs.Orderbook.Heartbook.Ask.Quoter = Quote.Quoter
				isBBO = true
				bboDesc = bboDescription{suit: hearts, quote: &Quote, direction: "ask"}
			}

		}
		for _, Quote := range cd.Bids {
			if Quote.Price > gs.Orderbook.Heartbook.Bid.Price {
				gs.Orderbook.Heartbook.Bid.Price = Quote.Price
				gs.Orderbook.Heartbook.Bid.Quoter = Quote.Quoter
				isBBO = true
				bboDesc = bboDescription{suit: hearts, quote: &Quote, direction: "bid"}
			}

		}

	case "diamonds":
		for _, Quote := range cd.Asks {
			if Quote.Price < gs.Orderbook.Heartbook.Ask.Price {
				gs.Orderbook.Diamondbook.Ask.Price = Quote.Price
				gs.Orderbook.Diamondbook.Ask.Quoter = Quote.Quoter
				isBBO = true
				bboDesc = bboDescription{suit: hearts, quote: &Quote, direction: "ask"}

			}

		}
		for _, Quote := range cd.Bids {
			if Quote.Quoter == myplayerName {
				continue
			} else if Quote.Price > gs.Orderbook.Diamondbook.Bid.Price {
				gs.Orderbook.Diamondbook.Bid.Price = Quote.Price
				gs.Orderbook.Diamondbook.Bid.Quoter = Quote.Quoter
				isBBO = true
				bboDesc = bboDescription{suit: hearts, quote: &Quote, direction: "bid"}

			}

		}

	}
	if isBBO {
		select {
		case bboChannel <- bboDesc:
		default:
		}
	}

}

func clearBook(suit Suit, gs *GameState, price int) {
	gs.rwMutex.Lock()
	defer gs.rwMutex.Unlock()
	//does lbo operation, sets it back to default
	gs.Orderbook.Clubbook.Ask.Price = 100
	gs.Orderbook.Diamondbook.Ask.Price = 100
	gs.Orderbook.Heartbook.Ask.Price = 100
	gs.Orderbook.Spadebook.Ask.Price = 100
	gs.Orderbook.Clubbook.Bid.Price = 0
	gs.Orderbook.Diamondbook.Bid.Price = 0
	gs.Orderbook.Heartbook.Bid.Price = 0
	gs.Orderbook.Spadebook.Bid.Price = 0
	switch suit {
	case spades:
		gs.Orderbook.Spadebook.LastPrice = price
	case clubs:
		gs.Orderbook.Clubbook.LastPrice = price
	case diamonds:
		gs.Orderbook.Diamondbook.LastPrice = price
	case hearts:
		gs.Orderbook.Heartbook.LastPrice = price
	}

	lboChannel <- struct{}{}

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
