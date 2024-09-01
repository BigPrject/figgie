package main

import "fmt"

// listens on the bbo
func listenOrder(client *Client, gs *GameState) {
	for {
		select {
		case <-ctx.Done():
			return
		case bbo := <-bboChannel:
			if bbo.quote.Quoter == myplayerName {
				continue
			} else {
				valueAlgo(bbo, client)

			}
		case <-lboChannel:
			arbAlgo(client, gs)
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

func valueAlgo(bbo bboDescription, client *Client) {
	switch bbo.direction {
	case "ask":
		if bbo.quote.Price < 5 {
			sendOrder(bbo.suit, "buy", bbo.quote.Price, client)
		}
		//possible conflict with fundamental, arbatariy values right now
	case "bid":
		if bbo.quote.Price > 12 {
			sendOrder(bbo.suit, "sell", bbo.quote.Price, client)
		}
	}
}

func arbAlgo(client *Client, gs *GameState) {
	card := gs.Trades[len(gs.Trades)-1].Card
	price := gs.Trades[len(gs.Trades)-1].Price / 2
	buyer := gs.Trades[len(gs.Trades)-1].Buyer

	if card != gs.goalSuit {
		if buyer == myplayerName && price == gs.arbPrice {
			//if the most recent trade was my arb buy I imdeatlly try to sell it and scalp. can change the threshold
			// it can be proven that my arb algo's price is unique compared to the bayesPrice and fundamentalBot, so I only need to check the price.
			sendOrder(card, "sell", price+2, client)
		} else {
			if price > 5 {
				sendOrder(card, "buy", 5, client)
			} else {
				sendOrder(card, "buy", price, client)
			}
			gs.arbPrice = price

		}
	}
}

// what if.... run once on every round

func exploitAlgo(client *Client) {
	order := Order{Card: "spade", Price: 1, Direction: "buy"}
	err := client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	order = Order{Card: "club", Price: 1, Direction: "buy"}
	err = client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	order = Order{Card: "heart", Price: 1, Direction: "buy"}
	err = client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	order = Order{Card: "diamond", Price: 1, Direction: "buy"}
	err = client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	order = Order{Card: "spade", Price: 99, Direction: "sell"}
	err = client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	order = Order{Card: "club", Price: 99, Direction: "sell"}
	err = client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	order = Order{Card: "heart", Price: 99, Direction: "sell"}
	err = client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	order = Order{Card: "diamond", Price: 99, Direction: "sell"}
	err = client.PlaceOrder(&order)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

}
