package bahirehasab

type BahireHasab struct {
	Year int
}

func (bh *BahireHasab) Medeb() int {
	year := bh.Year
	if (year+5500)%19 != 0 {
		return (year + 5500) % 19
	} else {
		return 19
	}
}

func (bh *BahireHasab) Wenber() int {
	return bh.Medeb() - 1
}

