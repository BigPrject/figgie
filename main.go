package main

func main() {
	client := newClient(wsURL, resURL, playerID)
	client.ConnectWebSocket()
	gs := NewGameState()
	//messageChan := make(chan []byte)
	// evntaully just have one function, start bot, that calls that go routine...
	go client.ListenToMessages(gs)

	//Run bayes then if no > 50% then start fund bot.
	go startFund(gs)

	gs.Inventory = &Inventory{
		Spades:   3,
		Clubs:    5,
		Diamonds: 2,
		Hearts:   0,
	}

	gs.Inventory.complexPrior()
}
