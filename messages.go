package main

import (
	"encoding/json"
	"fmt"
)

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
	Asks      []Qoute `json:"asks"`
	Bids      []Qoute `json:"bids"`
	LastTrade string  `json:"last_trade"`
}

func dealtCards(message Message) *Inventory {
	var inv *Inventory
	err := json.Unmarshal(message.Data, inv)
	if err != nil {
		fmt.Printf("Can't unmarhsall inventory %v", err)

		return nil
	}

	return inv
}

func endRound(message Message) {
	var end endRoundStuct
	err := json.Unmarshal(message.Data, &end)
	if err != nil {
		fmt.Printf("Can't unmarhsall round %v", err)

	}
	prettyPrintEndRound(end)
}

func endGame(message Message) {
	var end endRoundStuct
	err := json.Unmarshal(message.Data, &end)
	if err != nil {
		fmt.Printf("Can't unmarhsall round %v", err)

	}

}

func updateMessage(message Message) {
	var update UpdateStruct

	err := json.Unmarshal(message.Data, &update)
	if err != nil {
		fmt.Printf("Can't process update data %v", err)
	}

	processUpdate(update)
}
