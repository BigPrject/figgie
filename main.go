package main

var gs *GameState

func main() {
	client := newClient(wsURL, resURL, playerID)
	client.ConnectWebSocket()
	gs = NewGameState()
	messageChan := make(chan []byte)

	go client.ListenToMessages(messageChan)

	gs.Inventory = &Inventory{
		Spades:   3,
		Clubs:    5,
		Diamonds: 2,
		Hearts:   0,
	}

	gs.Inventory.complexPrior()
}
