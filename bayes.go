package main

import "fmt"

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
	/*
		goal suit can have 8 cards but we'll use ten
		so we don't mistakenlly overestimate
	*/

)

func (inv *Inventory) calcPrior() {
	cards := map[Suit]int{
		spade:   inv.spades,
		club:    inv.clubs,
		diamond: inv.diamonds,
		heart:   inv.hearts,
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
