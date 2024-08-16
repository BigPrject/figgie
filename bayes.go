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
)

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

func (inv *Inventory) complexPrior() {
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
	for card, prob := range complexProbs {
		fmt.Printf("%s had probability of %.2f\n", card, prob)
	}

	// Implement bayes that condsiders all other card in hand

}
