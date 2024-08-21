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
	inv         *Inventory
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

	go func(gs *GameState, fd *Fundbot) {
		for {
			select {
			case <-tradeChannel:
				fd.runFundamental(gs, client)
			}
		}
	}(gs, bot)
}

func (fd *Fundbot) runFundamental(gs *GameState, client *Client) {
	//Optimize later by making inventory be a map.
	hand := map[Suit]int{spades: fd.inv.Spades,
		clubs:    fd.inv.Clubs,
		diamonds: fd.inv.Diamonds, hearts: fd.inv.Hearts}
	fd.cardCount(gs.Trades)
	sums := fd.sumList()
	deckDistrbution := fd.calcMultinomal(sums)

	for suit, amount := range hand {
		expBuy := fd.expectedBuy(suit, amount, deckDistrbution[:])
		expSell := fd.expectedSell(suit, amount, deckDistrbution[:])
		sendOrder(suit, "buy", expBuy, client)
		sendOrder(suit, "buy", expSell, client)
		//send to order exec.
	}

}
