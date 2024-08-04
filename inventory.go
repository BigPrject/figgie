package main

type Inventory struct {
	spades   int
	clubs    int
	diamonds int
	hearts   int
}

func (inv *Inventory) update(s Suit, add bool) {
	switch s {
	case spade:
		if add {
			inv.spades += 1
		} else {
			inv.spades -= 1
		}
	case club:
		if add {
			inv.clubs += 1
		} else {
			inv.clubs -= 1
		}
	case heart:
		if add {
			inv.hearts += 1
		} else {
			inv.hearts -= 1
		}
	case diamond:
		if add {
			inv.diamonds += 1
		} else {
			inv.diamonds -= 1
		}

	}

}
