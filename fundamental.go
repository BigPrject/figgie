package main

import (
	"math"
)

// in the form {spades,clubs,hearts,diamonds,goalIndex}
/*
var suitNums = map[Suit]int{spades: 0, clubs: 1, hearts: 2, diamonds: 3}
var majorites = []int{5, 6, 6, 5, 6, 6, 6, 6, 5, 6, 6, 5}
var payouts = []int{120, 100, 100, 120, 100, 100, 100, 100, 120, 100, 100, 120}
var decks = [][]int{
	{12, 8, 10, 10, 1}, {12, 10, 8, 10, 1},
	{12, 10, 10, 8, 1}, {8, 12, 10, 10, 0},
	{10, 12, 8, 10, 0}, {10, 12, 10, 8, 0},
	{8, 10, 12, 10, 2}, {10, 8, 12, 10, 2},
	{10, 10, 12, 8, 2}, {8, 10, 10, 12, 3},
	{10, 8, 10, 12, 3}, {10, 8, 10, 12, 3},
}
var spadeList [4]int
var clubList [4]int
var heartList [4]int
var diamondList [4]int

func startList(inv *Inventory) {
	spadeList = [4]int{inv.Spades, 0, 0, 0}
	clubList = [4]int{inv.Clubs, 0, 0, 0}
	heartList = [4]int{inv.Hearts, 0, 0, 0}
	diamondList = [4]int{inv.Diamonds, 0, 0, 0}

}

*/

// card counting logic
func (fd *Fundbot) cardCount(gs *GameState) {
	for _, trade := range gs.Trades {
		switch trade.Card {
		case spades:
			cardAlgo(&fd.spadeList, &trade, gs)

		case clubs:
			cardAlgo(&fd.clubList, &trade, gs)
		case hearts:
			cardAlgo(&fd.heartList, &trade, gs)
		case diamonds:
			cardAlgo(&fd.diamondList, &trade, gs)
		}
	}

}

var playerIndex = 0

func cardAlgo(List *[4]int, trade *Trade, gs *GameState) {
	buyer, buyerPresent := gs.players[trade.Buyer]
	if !buyerPresent {
		gs.players[trade.Buyer] = (playerIndex % 4)
		buyer = playerIndex
		playerIndex++
	}
	seller, sellerPresent := gs.players[trade.Seller]
	if !sellerPresent {
		gs.players[trade.Seller] = (playerIndex % 4)
		seller = playerIndex
		playerIndex++
	}

	if List[seller] < 1 {
		List[buyer]++
		List[seller] = 0

	} else {
		List[buyer]++
		List[seller]--
	}

}

// iterate through all list and get sums

func (fd *Fundbot) sumList() []int {
	spadeSum := 0
	clubSum := 0
	heartSum := 0
	diamondSum := 0

	for i := 0; i < 4; i++ {
		spadeSum += fd.spadeList[i]
		clubSum += fd.clubList[i]
		heartSum += fd.heartList[i]
		diamondSum += fd.diamondList[i]

	}
	sum := []int{spadeSum, clubSum, heartSum, diamondSum}
	return sum
}

func (fd *Fundbot) calcMultinomal(sums []int) [12]float32 {
	var multinom [12]float32

	// probability of being in each a deck given the seen cards
	spade := sums[0]
	club := sums[1]
	heart := sums[2]
	diamond := sums[3]
	totalCombinations := 0

	combinations := make([]int, 12)

	for i := 0; i < 12; i++ {
		spadeNum := fd.decks[i][0]
		clubNum := fd.decks[i][1]
		heartNum := fd.decks[i][2]
		diamondNum := fd.decks[i][3]

		combPossible := combination(spadeNum, spade) * combination(clubNum, club) * combination(heartNum, heart) * combination(diamondNum, diamond)

		combinations[i] = combPossible
		totalCombinations += combPossible

	}

	for i := 0; i < 12; i++ {
		if totalCombinations == 0 {
			multinom[i] = 0
		} else {
			multinom[i] = float32(combinations[i]) / float32(totalCombinations)
		}

	}

	return multinom

}

func (fd *Fundbot) value(deckIndex int, card Suit, cardAmount int) int {
	goalSuit := fd.decks[deckIndex][4]
	curSuit := fd.suitNums[card]

	if goalSuit == curSuit || goalSuit == curSuit+1 || goalSuit == curSuit+2 {

		return 10 + fd.valuePayout(deckIndex, cardAmount)
	}
	return 0

}

func (fd *Fundbot) valuePayout(index int, amount int) int {
	const r = 4.5

	if amount < fd.majorites[index] {
		payout := fd.payouts[index]
		value := int(float64(payout) * (1 - r) * math.Pow(r, float64(amount)))
		return value
	}
	return 0
}

func (fd *Fundbot) expectedBuy(suit Suit, cards int, distrubiton [12]float32) int {
	expectedValue := 0
	// cards is amount of cards in my hand
	for i := 0; i < 12; i++ {
		expectedValue += int(distrubiton[i]) * fd.value(i, suit, cards)

	}
	return expectedValue

}

func (fd *Fundbot) expectedSell(suit Suit, cards int, distrubiton [12]float32) int {

	return fd.expectedBuy(suit, cards-1, distrubiton)

}
