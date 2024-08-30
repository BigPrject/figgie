package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	/*

			client := newClient(wsURL, resURL, playerID)
			client.ConnectWebSocket()
			gs := NewGameState()
			//send orders on start buys at 1, sells at 98, what if someone has a bug in their algo lol
			//messageChan := make(chan []byte)
			// evntaully just have one function, start bot, that calls that go routine...
			go client.ListenToMessages(gs)
			//*

			select {
			case <-invChannel:
				card, prob := gs.Inventory.complexPrior()
				if prob >= 0.5 {
					//run bayes bot
					go bayesBot(card, prob)
				} else {
					go startFund(gs, client)
				}
			}
				gs.Inventory = &Inventory{
			Spades:   7,
			Clubs:    1,
			Diamonds: 1,
			Hearts:   1,
		}
		gs.Inventory.englandCalc()
	*/
	//Run bayes then if no > 50% then start fund bot.

	var i int
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("How many rounds will we do: ")
	fmt.Scan(&i)
	startSimulation(i)
}
