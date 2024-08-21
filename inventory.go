package main

type Inventory struct {
	Spades   int `json:"spades"`
	Clubs    int `json:"clubs"`
	Diamonds int `json:"diamonds"`
	Hearts   int `json:"hearts"`
}

func (inv *Inventory) update(s Suit, add bool) {
	switch s {
	case spades:
		if add {
			inv.Spades += 1
		} else {
			inv.Spades -= 1
		}
	case clubs:
		if add {
			inv.Clubs += 1
		} else {
			inv.Clubs -= 1
		}
	case hearts:
		if add {
			inv.Hearts += 1
		} else {
			inv.Hearts -= 1
		}
	case diamonds:
		if add {
			inv.Diamonds += 1
		} else {
			inv.Diamonds -= 1
		}

	}

}
