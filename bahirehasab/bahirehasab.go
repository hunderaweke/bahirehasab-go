package bahirehasab

type BahireHasab struct {
	Year int
}

const (
	TINTEABEKTE = 11
	TINTEMETK   = 19
)

var (
	WERAT = []string{
		"መስከረም",
		"ጥቅምት",
		"ኅዳር",
		"ታኅሣስ",
		"ጥር",
		"የካቲት",
		"መጋቢት",
		"ሚያዝያ",
		"ግንቦት",
		"ሰኔ",
		"ሐምሌ",
		"ነሐሴ",
		"ጷግሜ",
	}
	WENGELAWI     = []string{"ዮሐንስ", "ማቴዎስ", "ማርቆስ", "ሉቃስ"}
	ELETAT        = []string{"ቅዳሜ", "እሑድ", "ሰኞ", "ማክሰኞ", "ረቡዕ", "ሐሙስ", "አርብ"}
	BEALAT_TEWSAK = []int{0, 14, 41, 62, 67, 69, 93, 108, 118, 119, 121}
	BEALAT        = []string{
		"ጾመ ነነዌ",
		"ዐብይ ጾም",
		"ደብረ ዘይት",
		"ሆሳዕና",
		"ስቅለት",
		"ትንሳኤ",
		"ርክበ ካሕናት",
		"ዕርገት",
		"በዓለ ሀምሳ",
		"ጾመ ሐዋርያት",
		"ጾመ ድኅነት",
	}
	ELETE_KEN   = []string{"ረቡዕ", "ሐሙስ", "አርብ", "ቅዳሜ", "እሑድ", "ሰኞ", "ማክሰኞ"}
	ELET_TEWSAK = map[string]int{"ሰኞ": 6, "ማክሰኞ": 5, "ረቡዕ": 4, "ሐሙስ": 3, "አርብ": 2, "ቅዳሜ": 8, "እሑድ": 7}
)

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
func (bh *BahireHasab) Abekte() int {
	if bh.Wenber()*TINTEABEKTE > 0 {
		return bh.Wenber() * TINTEABEKTE % 30
	} else {
		return 30
	}
}

func (bh *BahireHasab) Metk() int {
	if bh.Wenber()*TINTEMETK > 0 {
		return bh.Wenber() * TINTEMETK % 30
	} else {
		return 30
	}
}
