package main

type Fundbot struct {
	suitNums    map[Suit]int
	majorites   []int
	payouts     []int
	decks       [][]int
	spadeList   [4]int
	clubList    [4]int
	heartList   [4]int
	diamondList [4]int
}

func NewFundbot() *Fundbot {
	return &Fundbot{
		suitNums: map[Suit]int{
			spades:   0,
			clubs:    1,
			hearts:   2,
			diamonds: 3,
		},
		majorites: []int{5, 6, 6, 5, 6, 6, 6, 6, 5, 6, 6, 5},
		payouts:   []int{120, 100, 100, 120, 100, 100, 100, 100, 120, 100, 100, 120},
		decks: [][]int{
			{12, 8, 10, 10, 1},
			{12, 10, 8, 10, 1},
			{12, 10, 10, 8, 1},
			{8, 12, 10, 10, 0},
			{10, 12, 8, 10, 0},
			{10, 12, 10, 8, 0},
			{8, 10, 12, 10, 2},
			{10, 8, 12, 10, 2},
			{10, 10, 12, 8, 2},
			{8, 10, 10, 12, 3},
			{10, 8, 10, 12, 3},
			{10, 8, 10, 12, 3},
		},
	}
}

func startFund(gs *GameState, client *Client) {
	bot := NewFundbot()
	for {
		select {
		case <-ctx.Done():
			return
		case <-tradeChannel:
			bot.runFundamental(gs, client)
		}

	}

}

func (fd *Fundbot) runFundamental(gs *GameState, client *Client) {
	//Optimize later by making inventory be a map.

	fd.cardCount(gs)
	sums := fd.sumList()
	deckDistrbution := fd.calcMultinomal(sums)
	hand := getHand(gs)
	for suit, amount := range *hand {
		expBuy := fd.expectedBuy(suit, amount, deckDistrbution)
		expSell := fd.expectedSell(suit, amount, deckDistrbution)
		sendOrder(suit, "buy", expBuy, client)
		sendOrder(suit, "buy", expSell, client)
		//send to order exec.
	}

}
