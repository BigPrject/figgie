package main

type Qoute struct {
	price  int
	qouter string
}

type Orderbook struct {
	ask Qoute
	bid Qoute

	lastPrice int
}

func (book *Orderbook) newlastPrice(p int) {
	book.lastPrice = p
}

func newBook() Orderbook {

	return Orderbook{
		ask: Qoute{
			price:  0,
			qouter: "",
		},
		bid: Qoute{
			price:  99,
			qouter: "",
		},
		lastPrice: 0,
	}
}
