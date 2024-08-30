package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var Mydecks = [][]int{
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
}

func chooseDeck() int {
	return rand.IntN(12)
}

func makeInv(deckIdx int) (*Inventory, Suit) {
	cardsToGo := 10
	var commonSuit Suit
	var spadesAmount int
	var clubsAmount int
	var heartsAmount int
	var diamondsAmount int

	if Mydecks[deckIdx][0] == 12 {
		commonSuit = spades
	} else if Mydecks[deckIdx][1] == 12 {
		commonSuit = clubs

	} else if Mydecks[deckIdx][2] == 12 {
		commonSuit = hearts
	} else if Mydecks[deckIdx][3] == 12 {
		commonSuit = diamonds
	}

	spadesAmount = rand.IntN(cardsToGo + 1)
	cardsToGo -= spadesAmount
	clubsAmount = rand.IntN(cardsToGo + 1)
	cardsToGo -= clubsAmount
	heartsAmount = rand.IntN(cardsToGo + 1)
	cardsToGo -= heartsAmount
	diamondsAmount = rand.IntN(cardsToGo + 1)
	cardsToGo -= diamondsAmount

	if cardsToGo != 0 {
		suit := rand.IntN(4)

		switch suit {
		case 0:
			spadesAmount += cardsToGo
		case 1:
			clubsAmount += cardsToGo
		case 2:
			heartsAmount += cardsToGo
		case 3:
			diamondsAmount += cardsToGo

		}
	}

	inv := &Inventory{Spades: spadesAmount, Clubs: clubsAmount,
		Hearts: heartsAmount, Diamonds: diamondsAmount}
	return inv, commonSuit
}

func startSimulation(rounds int) {
	trial := 0
	EnglandCorrect := 0
	NormalCorrect := 0
	ComplexCorrect := 0

	for trial < rounds {

		inv, common := makeInv(chooseDeck())

		cardEngland, amountEngland := inv.englandCalc()
		cardNormal, amountNormal := inv.calcPrior()
		cardComplex, amountComplex := inv.complexPrior()

		// could justhave the stuff return the card
		if amountEngland >= .50 {
			if cardEngland == common {
				EnglandCorrect++
			} else {
				EnglandCorrect--
			}
		}
		if amountNormal > .50 {
			if cardNormal == common {
				NormalCorrect++
			} else {
				NormalCorrect--
			}
		}
		if amountComplex >= .65 {
			if cardComplex == common {
				ComplexCorrect++
			} else {
				ComplexCorrect--
			}
		}

		trial++

		//fmt.Printf("Inventory | Spades: %d | Clubs %d | Hearts %d | Diamonds %d || Common Suit: %s\n", inv.Spades, inv.Clubs, inv.Hearts, inv.Diamonds, common)

	}
	printWins(EnglandCorrect, NormalCorrect, ComplexCorrect, trial)
}

func printWins(eng int, norm int, complex int, trial int) {

	timestamp := time.Now().Format("15:04:05.000")
	fmt.Printf("[%s] [Trial %d] England Correct : %d\t Complex Correct : %d\t Normal Correct : %d\n", timestamp, trial, eng, complex, norm)
}
