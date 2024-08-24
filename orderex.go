package main

import "fmt"

// listens on the bbo
func listenOrder(gs *GameState) {
	for {
		select {
		case bbo := <-bookChannel:
			if bbo.quote.Quoter == myplayerName {
				continue
			} else {
				// call bayes algo to determine wether to buy or sell...
			}
		}
	}

}

func sendOrder(suit Suit, direction string, amount int, client *Client) {
	var card string
	switch suit {
	case spades:
		card = "spade"
	case clubs:
		card = "club"
	case hearts:
		card = "heart"
	case diamonds:
		card = "diamond"

	}
	// impelent sanity check here, i.e max price I'd be willing to sell at etc etc,
	order := Order{Card: card, Price: amount, Direction: direction}

	err := client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

}

// what if.... run once on every round
func expliotAlgo(client *Client) {
	order := Order{Card: "spade", Price: 1, Direction: "buy"}
	err := client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	order = Order{Card: "club", Price: 1, Direction: "buy"}
	err = client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	order = Order{Card: "heart", Price: 1, Direction: "buy"}
	err = client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	order = Order{Card: "diamond", Price: 1, Direction: "buy"}
	err = client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	order = Order{Card: "spade", Price: 99, Direction: "sell"}
	err = client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	order = Order{Card: "club", Price: 99, Direction: "sell"}
	err = client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	order = Order{Card: "heart", Price: 99, Direction: "sell"}
	err = client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	order = Order{Card: "diamond", Price: 99, Direction: "sell"}
	err = client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

}
