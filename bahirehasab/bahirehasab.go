package bahirehasab

import (
	"slices"
	"strconv"
	"strings"
)

type BahireHasab struct {
	Year int
}

const (
	TINTEABEKTE = 11
	TINTEMETK   = 19
)

var (
	WERAT = []string{"መስከረም",
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
		"ጷግሜ"}
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
	ELET_TEWSAK = map[string]int{
		"ሰኞ":   6,
		"ማክሰኞ": 5,
		"ረቡዕ":  4,
		"ሐሙስ":  3,
		"አርብ":  2,
		"ቅዳሜ":  8,
		"እሑድ":  7}
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

func (bh *BahireHasab) Wengelawi() string {
	return WENGELAWI[(bh.Year+5500)%4]
}
func (bh *BahireHasab) MeteneRabiet() int {
	return int((bh.Year + 5500) / 4)
}

func (bh *BahireHasab) BealeMetk() string {
	var bealeMetk string
	if 15 <= bh.Metk() && bh.Metk() <= 30 {
		bealeMetk = "መስከረም " + strconv.Itoa(bh.Metk())
	} else if 2 <= bh.Metk() && bh.Metk() <= 14 {
		bealeMetk = "ጥቀምት " + strconv.Itoa(bh.Metk())
	}
	return bealeMetk
}
func (bh *BahireHasab) EletKen(e string) string {
	elet := strings.Split(e, " ")
	atsfeWer := (slices.Index(WERAT, elet[0]) + 1) * 2
	tnteYon := (bh.MeteneRabiet()+bh.Year+5500)%7 - 1
	kenStr := elet[len(elet)-1]
	ken, err := strconv.Atoi(kenStr)
	if err != nil {
		panic(err)
	}
	return ELETAT[(ken+tnteYon+atsfeWer)%7]
}
func (bh *BahireHasab) NewYear() string {
	ameteAlem := bh.Year + 5500
	a := (ameteAlem + bh.MeteneRabiet() + 2) % 7
	return ELETAT[a]

}

func (bh *BahireHasab) MebajaHamer() int {
	_mebajaHamer := bh.Metk() + ELET_TEWSAK[bh.EletKen(bh.BealeMetk())]
	return _mebajaHamer
}

func (bh *BahireHasab) Neneweh() string {
	_bealeMetk := bh.BealeMetk()
	_l := strings.Split(_bealeMetk, " ")
	_mebajaHamer := bh.MebajaHamer()
	var _wer string
	if _mebajaHamer > 30 {
		_wer = "የካቲት"
		_mebajaHamer %= 30
	} else if bh.Metk() == 30 || bh.Metk() == 0 {
		_wer = "የካቲት"
	} else if _l[0] == "መስከረም" {
		_wer = "ጥር"
	} else {
		_wer = "የካቲት"
	}
	return _wer + " " + strconv.Itoa(_mebajaHamer)
}

func (bh *BahireHasab) AtswamatWebealat(beal string) string {
	_bealTewsak := BEALAT_TEWSAK[slices.Index(BEALAT, beal)]
	if _bealTewsak == 0 {
		return bh.Neneweh()
	}
	_neneweh := strings.Split(bh.Neneweh(), " ")
	_nenewehWer, _nenewehKenStr := _neneweh[0], _neneweh[1]
	_nenewehKen, err := strconv.Atoi(_nenewehKenStr)
	if err != nil {
		panic(err)
	}
	_bealeKen := _nenewehKen + _bealTewsak
	var _bealeWer string
	if _bealeKen%30 == 0 {
		_bealeWer = WERAT[slices.Index(WERAT, _nenewehWer)+(_bealeKen/30)-1]
	} else {
		_bealeWer = WERAT[slices.Index(WERAT, _nenewehWer)+(_bealeKen/30)]
	}
	if _bealeKen%30 != 0 {
		_bealeKen %= 30
	} else {
		_bealeKen = 30
	}
	return _bealeWer + " " + strconv.Itoa(_bealeKen)
}

func (bh *BahireHasab) AbiyTsom() string {
	return bh.AtswamatWebealat(BEALAT[1])
}
func (bh *BahireHasab) DebreZeyt() string {
	return bh.AtswamatWebealat(BEALAT[2])
}
func (bh *BahireHasab) Hosaena() string {
	return bh.AtswamatWebealat(BEALAT[3])
}
func (bh *BahireHasab) Seklet() string {
	return bh.AtswamatWebealat(BEALAT[4])
}
func (bh *BahireHasab) Tnsae() string {
	return bh.AtswamatWebealat(BEALAT[5])
}
func (bh *BahireHasab) RkbeKahnat() string {
	return bh.AtswamatWebealat(BEALAT[6])
}
func (bh *BahireHasab) Erget() string {
	return bh.AtswamatWebealat(BEALAT[7])
}
func (bh *BahireHasab) BealeHamsa() string {
	return bh.AtswamatWebealat(BEALAT[8])
}

func (bh *BahireHasab) TsomeHawaryat() string {
	return bh.AtswamatWebealat(BEALAT[9])
}

func (bh *BahireHasab) TsomeDihnet() string {
	return bh.AtswamatWebealat(BEALAT[10])
}
