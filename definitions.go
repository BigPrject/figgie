package main

type Suit string

type Trade struct {
	card   Suit
	price  int
	buyer  string
	seller string
}

const (
	spade   Suit = "spades"
	club    Suit = "clubs"
	diamond Suit = "diamonds"
	heart   Suit = "hearts"
)

const (
	wsURL      = "ws://testnet-ws.figgiewars.com"
	resURL     = "http://testnet.figgiewars.com"
	playerID   = "LebronJames"
	playerName = "bellamy"
)

func (s Suit) getGoalSuit() Suit {
	switch s {
	case spade:
		return club
	case club:
		return spade
	case diamond:
		return heart
	case heart:
		return diamond
	}

	return ""
}
