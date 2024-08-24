package main

import (
	"encoding/json"
	"fmt"
)

type endGameStruct struct {
	PlayerPoints []playerPoints `json:"player_points"`
}
type endRoundStuct struct {
	CommonSuit      string `json:"common_suit"`
	GoalSuit        string `json:"goal_suit"`
	playerInventory []struct {
		PlayerName string `json:"player_name"`
		Spades     int    `json:"spades"`
		Clubs      int    `json:"clubs"`
		Diamonds   int    `json:"diamonds"`
		Hearts     int    `json:"hearts"`
	}
	PlayerPoints []playerPoints `json:"player_points"`
}

type playerPoints struct {
	PlayerName string `json:"player_name"`
	Points     int    `json:"points"`
}

type UpdateStruct struct {
	Spades   CardData `json:"spades"`
	Clubs    CardData `json:"clubs"`
	Diamonds CardData `json:"diamonds"`
	Hearts   CardData `json:"hearts"`
	Trade    string   `json:"trade"`
}

type CardData struct {
	Asks      []Quote `json:"asks"`
	Bids      []Quote `json:"bids"`
	LastTrade string  `json:"last_trade"`
}

func handleMessage(payload []byte, gs *GameState) {
	var m Message
	err := json.Unmarshal(payload, &m)
	if err != nil {
		fmt.Printf("Couldn't unmarshall paylaod %v", err)
	}
	switch m.Kind {
	case update:
		updateMessage(m, gs)
	case dealing:
		dealtCards(m, gs)
	case endOfGame:
		endGame(m, gs)
	case endOfRound:
		endRound(m, gs)
	}

}

func dealtCards(message Message, gs *GameState) {
	var inv *Inventory
	err := json.Unmarshal(message.Data, inv)
	if err != nil {
		fmt.Printf("Can't unmarhsall inventory %v", err)
	}
	gs.Inventory = inv
	gs.Balance -= 50
	select {
	case invChannel <- struct{}{}:
	default:
	}
}

func endRound(message Message, gs *GameState) {
	var end endRoundStuct
	err := json.Unmarshal(message.Data, &end)
	if err != nil {
		fmt.Printf("Can't unmarhsall round %v", err)

	}
	prettyPrintEndRound(end)

	gs.Inventory = &Inventory{}
	gs.Orderbook = NewOrderbook()
	gs.Trades = make([]Trade, 0)

}

func endGame(message Message, gs *GameState) {
	var end endGameStruct
	err := json.Unmarshal(message.Data, &end)
	if err != nil {
		fmt.Printf("Can't unmarhsall round %v", err)

	}
	prettyPrintEndGame(end)

	//reset my gamestate, shorten it to newgamestate later..
	gs.Inventory = &Inventory{}
	gs.Orderbook = NewOrderbook()
	gs.Trades = make([]Trade, 100)
	gs.Balance = 350

}

func updateMessage(message Message, gs *GameState) {
	var update UpdateStruct

	err := json.Unmarshal(message.Data, &update)
	if err != nil {
		fmt.Printf("Can't process update data %v", err)
	}

	processUpdate(update, gs)
}
