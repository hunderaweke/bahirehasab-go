package bahirehasab

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

type BahireHasab struct {
	Year   int
	Logger log.Logger
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

func (bh *BahireHasab) wenber() int {
	_wenber := bh.Medeb() - 1
	bh.Logger.Printf("wenber Returned: %v\n", _wenber)
	if _wenber == 0 {
		return 18
	}
	return _wenber
}
func (bh *BahireHasab) abekte() int {
	var _abekte int
	if bh.wenber()*TINTEABEKTE > 0 {
		_abekte = (bh.wenber() * TINTEABEKTE) % 30
	} else {
		_abekte = 30
	}
	bh.Logger.Printf("abekte returned: %v\n", _abekte)
	return _abekte
}

func (bh *BahireHasab) metk() int {
	var _metk int
	if bh.wenber()*TINTEMETK != 0 {
		_metk = (bh.wenber() * TINTEMETK) % 30
	} else {
		_metk = 30
	}
	bh.Logger.Printf("metk returned: %v\n", _metk)
	return _metk
}

func (bh *BahireHasab) Wengelawi() string {
	return WENGELAWI[(bh.Year+5500)%4]
}
func (bh *BahireHasab) meteneRabiet() int {
	bh.Logger.Printf("meteneRabiet returned: %v\n", int((bh.Year+5500)/4))
	return int((bh.Year + 5500) / 4)
}

func (bh *BahireHasab) bealeMetk() string {
	var bealeMetk string
	metk := bh.metk() % 30
	if 15 <= metk && metk <= 30 {
		bealeMetk = "መስከረም " + strconv.Itoa(metk)
	} else if 2 <= metk && metk <= 14 {
		bealeMetk = "ጥቅምት " + strconv.Itoa(metk)
	}
	bh.Logger.Printf("bealeMetk returned: %v\n", bealeMetk)
	return bealeMetk
}
func (bh *BahireHasab) EletKen(e string) string {
	bh.Logger.Printf("entering EletKen with: %v\n", e)
	elet := strings.Split(e, " ")
	bh.Logger.Printf("elet: %v\n", elet)
	atsfeWer := (slices.Index(WERAT, elet[0]) + 1) * 2
	bh.Logger.Printf("EletKen atsfewer: %v\n", atsfeWer)
	tnteYon := (bh.meteneRabiet()+bh.Year+5500)%7 - 1
	kenStr := elet[len(elet)-1]
	ken, err := strconv.Atoi(kenStr)
	if err != nil {
		bh.Logger.Printf("error in making: %v msg %v\n", ken, err)
	}
	return ELETAT[(ken+tnteYon+atsfeWer)%7]
}
func (bh *BahireHasab) NewYear() string {
	ameteAlem := bh.Year + 5500
	a := (ameteAlem + bh.meteneRabiet() + 2) % 7
	return ELETAT[a]

}

func (bh *BahireHasab) mebajaHamer() int {
	eletKen := bh.EletKen(bh.bealeMetk())
	_mebajaHamer := bh.metk() + ELET_TEWSAK[eletKen]
	bh.Logger.Printf("getting eletken: %v\n", eletKen)
	bh.Logger.Printf("getting elet_tewsak as : %v\n", ELET_TEWSAK[bh.EletKen(bh.bealeMetk())])
	bh.Logger.Printf("mebajaHamer returned: %v\n", _mebajaHamer)
	return _mebajaHamer
}

func (bh *BahireHasab) Neneweh() string {
	_bealeMetk := bh.bealeMetk()
	_l := strings.Split(_bealeMetk, " ")
	_mebajaHamer := bh.mebajaHamer()
	var _wer string
	if _mebajaHamer > 30 {
		_wer = "የካቲት"
		_mebajaHamer %= 30
	} else if bh.metk() == 30 || bh.metk() == 0 {
		_wer = "የካቲት"
	} else if _l[0] == "መስከረም" {
		_wer = "ጥር"
	} else {
		_wer = "የካቲት"
	}
	return _wer + " " + strconv.Itoa(_mebajaHamer)
}

func (bh *BahireHasab) atswamatWebealat(beal string) (string, string) {
	bh.Logger.Printf("entering AtswamatWebealat with: %v\n", beal)
	_bealTewsak := BEALAT_TEWSAK[slices.Index(BEALAT, beal)]
	if _bealTewsak == 0 {
		return bh.Neneweh(), beal
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
	return beal, (_bealeWer + " " + strconv.Itoa(_bealeKen))
}

func (bh *BahireHasab) AbiyTsom() (string, string) {
	return bh.atswamatWebealat(BEALAT[1])
}
func (bh *BahireHasab) DebreZeyt() (string, string) {
	return bh.atswamatWebealat(BEALAT[2])
}
func (bh *BahireHasab) Hosaena() (string, string) {
	return bh.atswamatWebealat(BEALAT[3])
}
func (bh *BahireHasab) Seklet() (string, string) {
	return bh.atswamatWebealat(BEALAT[4])
}
func (bh *BahireHasab) Tnsae() (string, string) {
	return bh.atswamatWebealat(BEALAT[5])
}
func (bh *BahireHasab) RkbeKahnat() (string, string) {
	return bh.atswamatWebealat(BEALAT[6])
}
func (bh *BahireHasab) Erget() (string, string) {
	return bh.atswamatWebealat(BEALAT[7])
}
func (bh *BahireHasab) BealeHamsa() (string, string) {
	return bh.atswamatWebealat(BEALAT[8])
}

func (bh *BahireHasab) TsomeHawaryat() (string, string) {
	return bh.atswamatWebealat(BEALAT[9])
}

func (bh *BahireHasab) TsomeDihnet() (string, string) {
	return bh.atswamatWebealat(BEALAT[10])
}
