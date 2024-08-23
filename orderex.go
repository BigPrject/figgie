package main

import "fmt"

func listenOrder(gs *GameState) {
	for {
		select {
		case <-bookChannel:
			// do action when update on bbo.
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
