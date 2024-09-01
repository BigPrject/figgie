package main

import (
	"context"
	"fmt"
)

func main() {

	ctx, cancelCtx = context.WithCancel(context.Background())

	client := newClient(wsURL, resURL, playerID)
	if err := client.registerTestNet(playerID); err != nil {
		fmt.Printf("Error registering: %v\n", err)
		return
	}

	if err := client.ConnectWebSocket(); err != nil {
		fmt.Printf("Error connecting WebSocket: %v\n", err)
		return
	}
	gs := NewGameState()

	wg.Add(1)
	go func() {
		defer wg.Done()
		client.ListenToMessages(gs)
	}()

	wg.Wait()
	//send orders on start buys at 1, sells at 98, what if someone has a bug in their algo lol
	//messageChan := make(chan []byte)
	// evntaully just have one function, start bot, that calls that go routine...

	//*

	//Run bayes then if no > 50% then start fund bot.
	/*
			gs.Inventory = &Inventory{
			Spades:   7,
			Clubs:    1,
			Diamonds: 1,
			Hearts:   1,
		}


			var i int
			fmt.Println("How many rounds will we do: ")
			fmt.Scan(&i)
			startSimulation(i)
	*/
}
