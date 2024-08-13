package main

var inventory *Inventory

func main() {
	myInv := Inventory{
		spades:   0,
		clubs:    6,
		diamonds: 4,
		hearts:   0,
	}

	myInv.complexPrior()
}
