package main

import (
	"fmt"
	"math"
)

const (
	cards              = 40
	suits              = 4
	initProb           = 0.25
	uniqueSuit         = 12
	goalSuit           = 10
	hand               = 10
	remainingNonUnique = 28
	remainingNonGoal   = 30

	totalHands = 847660528
)

func (inv *Inventory) calcPrior() {
	cards := map[Suit]int{
		spades:   inv.Spades,
		clubs:    inv.Clubs,
		diamonds: inv.Diamonds,
		hearts:   inv.Hearts,
	}

	probs := make(map[Suit]float32)

	for card, amount := range cards {
		restOfHand := hand - amount

		combOfCardCommon := float32(combination(uniqueSuit, amount)) * float32(combination(remainingNonUnique, restOfHand))
		combofCard10 := float32(combination(hand, amount)) * float32(combination(remainingNonGoal, restOfHand))

		probCardCommon := float32(combOfCardCommon / float32(totalHands))
		probCardNotCommon := float32(combofCard10 / float32(totalHands))
		probAmountSpades := (probCardCommon * initProb) + (probCardNotCommon * (1 - initProb))

		if probAmountSpades == 0 {
			fmt.Printf("Suit %s, probability is undefined due to zero denominator\n", card)
			continue
		}

		bayesCalc := probCardCommon * initProb / probAmountSpades
		probs[card] = bayesCalc
	}

	for suit, prob := range probs {
		fmt.Printf("Suit %s, probability %.2f\n", suit, prob)
	}
}

func (inv *Inventory) complexPrior() (Suit, float32) {
	cards := map[Suit]int{
		spades:   inv.Spades,
		clubs:    inv.Clubs,
		diamonds: inv.Diamonds,
		hearts:   inv.Hearts,
	}
	complexProbs := make(map[Suit]float32)

	for card, amount := range cards {
		otherCards := make(map[Suit]int)

		switch card {
		case spades, clubs:
			otherCards[diamonds] = inv.Diamonds
			otherCards[hearts] = inv.Hearts
		case diamonds, hearts:
			otherCards[spades] = inv.Spades
			otherCards[clubs] = inv.Clubs
		}
		for otherCard, count := range otherCards {
			restOfHand := hand - amount - count

			combCardCommonOtherCard := float32(combination(12, amount) * combination(10, count) * combination(18, restOfHand))

			combCardOtherCard := float32(combination(10, amount) * combination(10, count) * combination(20, restOfHand))

			probCardCommonOtherCard := combCardCommonOtherCard / totalHands
			probCardNotCommon := combCardOtherCard / totalHands

			probCardOtherCard := (probCardCommonOtherCard * initProb) + (probCardNotCommon * (1 - initProb))

			bayesComplex := (probCardCommonOtherCard * initProb) / probCardOtherCard

			// maybe use max function instead
			if complexProbs[card] == 0 {
				complexProbs[card] = bayesComplex
			} else {
				complexProbs[card] = max(bayesComplex, complexProbs[card])
			}

			fmt.Printf("Probability of %s given %d of %s is %.2f\n", card, count, otherCard, bayesComplex)

		}

	}
	fmt.Println("\nMAX PROBABILITES")
	max := float32(0)
	var maxCard Suit
	for card, prob := range complexProbs {
		fmt.Printf("%s had probability of %.2f\n", card, prob)

		if prob > max {
			max = prob
			maxCard = card
		}
	}
	return maxCard, max
	// Implement bayes that condsiders all other card in hand

}

func (inv *Inventory) englandCalc() {
	cards := map[Suit]int{
		spades:   inv.Spades,
		clubs:    inv.Clubs,
		diamonds: inv.Diamonds,
		hearts:   inv.Hearts,
	}

	numWaysSpadesCommonDeckA := combination(12, inv.Spades) * combination(10, inv.Clubs) * combination(10, inv.Hearts) * combination(8, inv.Diamonds)
	numWaysSpadesCommonDeckB := combination(12, inv.Spades) * combination(10, inv.Clubs) * combination(8, inv.Hearts) * combination(10, inv.Diamonds)
	numWaysSpadesCommonDeckC := combination(12, inv.Spades) * combination(8, inv.Clubs) * combination(10, inv.Hearts) * combination(10, inv.Diamonds)
	numWaysClubCommonDeckA := combination(12, inv.Clubs) * combination(10, inv.Spades) * combination(10, inv.Hearts) * combination(8, inv.Diamonds)
	numWaysClubCommonDeckB := combination(12, inv.Clubs) * combination(10, inv.Spades) * combination(8, inv.Hearts) * combination(10, inv.Diamonds)
	numWaysClubCommonDeckC := combination(12, inv.Clubs) * combination(8, inv.Spades) * combination(10, inv.Hearts) * combination(10, inv.Diamonds)
	numWaysHeartCommonDeckA := combination(12, inv.Hearts) * combination(10, inv.Spades) * combination(10, inv.Clubs) * combination(8, inv.Diamonds)
	numWaysHeartCommonDeckB := combination(12, inv.Hearts) * combination(10, inv.Spades) * combination(8, inv.Clubs) * combination(10, inv.Diamonds)
	numWaysHeartCommonDeckC := combination(12, inv.Hearts) * combination(8, inv.Spades) * combination(10, inv.Clubs) * combination(10, inv.Diamonds)
	numWaysDiamondCommonDeckA := combination(12, inv.Diamonds) * combination(10, inv.Spades) * combination(10, inv.Clubs) * combination(8, inv.Hearts)
	numWaysDiamondCommonDeckB := combination(12, inv.Diamonds) * combination(10, inv.Spades) * combination(8, inv.Clubs) * combination(10, inv.Hearts)
	numWaysDiamondCommonDeckC := combination(12, inv.Diamonds) * combination(8, inv.Spades) * combination(10, inv.Clubs) * combination(10, inv.Hearts)

	for card, _ := range cards {
		var totalWays float32
		var probBgivenA float32
		var bayescalc float32
		switch card {
		case spades:
			totalWays := (numWaysSpadesCommonDeckA * 1 / 3) + (numWaysSpadesCommonDeckB * 1 / 3) + (numWaysSpadesCommonDeckC * 1 / 3)
			probBgivenA := float32(totalWays / totalHands)
			probBnotA := (float32(1 / 3 * (numWaysClubCommonDeckA + numWaysClubCommonDeckB + numWaysClubCommonDeckC + numWaysHeartCommonDeckA + numWaysHeartCommonDeckB + numWaysHeartCommonDeckC + numWaysDiamondCommonDeckA + numWaysDiamondCommonDeckB + numWaysDiamondCommonDeckC))) / totalHands
			probB := (probBgivenA * initProb) + (probBnotA * float32(1-initProb))
			bayescalc := (float32(probBgivenA) * (initProb) / probB)

		}

	}
}

// refine this later
func bayesPrice(prior float32) int {
	maxPay := 22.0
	// stepness of curve
	k := 10.0
	mid := 0.5
	// add a sum that accounts for the amount I have in my hand, more I have in my hand, more I should be willing to pay.
	return int(maxPay / (1 + math.Exp(-k*(float64(prior)-mid))))

}

func bayesBot(card Suit, prob float32) {
	goalSuit := card.getGoalSuit()
	goalSuit.getGoalSuit()
	for {
		//listens on order and runs bayes bot
	}
}

func combination(n, k int) int {
	if k > n {
		return 0
	}
	if k > n/2 {
		k = n - k
	}
	comb := 1
	for i := 1; i <= k; i++ {
		comb = (n - k + i) * comb / i
	}
	return comb
}

// call this when bayes calc is valid , this should off all other suit' and gain goal suit
func bayesAlgo(card Suit, gs *GameState) {

}
