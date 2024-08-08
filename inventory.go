package main

type Inventory struct {
	spades   int `json: "spades"`
	clubs    int `json: "clubs"`
	diamonds int `json: "diamonds"`
	hearts   int `json: "hearts"`
}

func (inv *Inventory) update(s Suit, add bool) {
	switch s {
	case spades:
		if add {
			inv.spades += 1
		} else {
			inv.spades -= 1
		}
	case clubs:
		if add {
			inv.clubs += 1
		} else {
			inv.clubs -= 1
		}
	case hearts:
		if add {
			inv.hearts += 1
		} else {
			inv.hearts -= 1
		}
	case diamonds:
		if add {
			inv.diamonds += 1
		} else {
			inv.diamonds -= 1
		}

	}

}
