package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
)

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
	clearConsole()
}

func prettyPrintEndGame(data endGameStruct) {
	const darkGreen = "\033[32m"
	const reset = "\033[0m"

	pointsArr := make([]int, 4)
	nameToPoints := make(map[string]int)

	for i, player := range data.PlayerPoints {
		pointsArr[i] = player.Points
		nameToPoints[player.PlayerName] = player.Points
		prettyPrintPlayerPoints(player)

	}
	// compartor sorts in deceding order
	slices.SortFunc(pointsArr, func(a, b int) int {
		if a > b {
			return -1
		} else if b < a {
			return 1
		} else {
			return 0
		}
	})

	if nameToPoints[myplayerName] != pointsArr[0] || nameToPoints[myplayerName] != pointsArr[1] {
		fmt.Println(darkGreen + "You lost the game..." + reset)

	} else {
		fmt.Println(darkGreen + "You were top 2!!! moving on to the next game" + reset)

	}
	// if my name is not top 2 print fails

	clearConsole()
}
func prettyPrintPlayerPoints(pp playerPoints) {
	const purple = "\033[35m"
	const reset = "\033[0m"

	fmt.Println(purple + "Player Points:" + reset)
	fmt.Println(purple + "----------------" + reset)
	fmt.Printf(purple+"Player Name: %s\n"+reset, pp.PlayerName)
	fmt.Printf(purple+"Points: %d\n"+reset, pp.Points)
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
