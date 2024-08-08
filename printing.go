package main

import "fmt"

// Function to print the EndRoundStuct data in purple color
func prettyPrintEndRound(data endRoundStuct) {
	const purple = "\033[35m"
	const reset = "\033[0m"

	fmt.Println(purple + "End Round Data:" + reset)
	fmt.Println(purple + "----------------" + reset)

	fmt.Printf(purple+"Common Suit: %s\n"+reset, data.CommonSuit)
	fmt.Printf(purple+"Goal Suit: %s\n"+reset, data.GoalSuit)

	fmt.Println(purple + "Player Points:" + reset)
	for _, pt := range data.PlayerPoints {
		prettyPrintPlayerPoints(pt)
	}

	fmt.Println(purple + "Player Inventories:" + reset)
	for _, inv := range data.playerInventory {
		fmt.Printf(purple+"  Player Name: %s\n"+reset, inv.PlayerName)
		fmt.Printf(purple+"  Spades: %d\n"+reset, inv.Spades)
		fmt.Printf(purple+"  Clubs: %d\n"+reset, inv.Clubs)
		fmt.Printf(purple+"  Diamonds: %d\n"+reset, inv.Diamonds)
		fmt.Printf(purple+"  Hearts: %d\n"+reset, inv.Hearts)
		fmt.Println()
	}

}

func prettyPrintPlayerPoints(pp playerPoints) {
	const purple = "\033[35m"
	const reset = "\033[0m"

	fmt.Println(purple + "Player Points:" + reset)
	fmt.Println(purple + "----------------" + reset)
	fmt.Printf(purple+"Player Name: %s\n"+reset, pp.PlayerName)
	fmt.Printf(purple+"Points: %d\n"+reset, pp.Points)
}
