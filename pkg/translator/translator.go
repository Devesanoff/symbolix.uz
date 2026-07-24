package translator

import (
	"strings"
)


// alphabetMode parametri: "standard" (Oʻ, Gʻ, Sh, Ch) yoki "new" (Ö, Ğ, Ş, Ç) bo'lishi mumkin.
func CyrillicToLatin(text string, alphabetMode string) string {
	var specificReplacements []struct {
		cyr string
		lat string
	}

	// 1. Foydalanuvchi tanlagan rejimga qarab maxsus harflar ro'yxatini yuklaymiz
	if alphabetMode == "new" {
		// Yangi (taklif etilgan) alifbo
		specificReplacements = []struct{ cyr, lat string }{
			{"Ё", "Yo"}, {"ё", "yo"},
			{"Ю", "Yu"}, {"ю", "yu"},
			{"Я", "Ya"}, {"я", "ya"},
			{"Ч", "Ç"}, {"ч", "ç"},
			{"Ш", "Ş"}, {"ш", "ş"},
			{"Щ", "Ş"}, {"щ", "ş"},
			{"Ў", "Ö"}, {"ў", "ö"},
			{"Қ", "Q"}, {"қ", "q"},
			{"Ғ", "Ğ"}, {"ғ", "ğ"},
			{"Ҳ", "H"}, {"ҳ", "h"},
			{"Ц", "Ts"}, {"ц", "ts"},
			{"Ъ", "’"}, {"ъ", "’"},
			{"Ь", ""}, {"ь", ""},
		}
	} else {
		// Amaldagi standart alifbo
		specificReplacements = []struct{ cyr, lat string }{
			{"Ё", "Yo"}, {"ё", "yo"},
			{"Ю", "Yu"}, {"ю", "yu"},
			{"Я", "Ya"}, {"я", "ya"},
			{"Ч", "Ch"}, {"ч", "ch"},
			{"Ш", "Sh"}, {"ш", "sh"},
			{"Щ", "Sh"}, {"щ", "sh"},
			{"Ў", "Oʻ"}, {"ў", "oʻ"},
			{"Қ", "Q"}, {"қ", "q"},
			{"Ғ", "Gʻ"}, {"ғ", "gʻ"},
			{"Ҳ", "H"}, {"ҳ", "h"},
			{"Ц", "Ts"}, {"ц", "ts"},
			{"Ъ", "’"}, {"ъ", "’"},
			{"Ь", ""}, {"ь", ""},
		}
	}

	// 2. Ikkala alifbo uchun o'zgarmas (umumiy) bo'lgan asosiy harflar
	baseReplacements := []struct{ cyr, lat string }{
		{"А", "A"}, {"а", "a"}, {"Б", "B"}, {"б", "b"}, {"В", "V"}, {"в", "v"},
		{"Г", "G"}, {"г", "g"}, {"Д", "D"}, {"д", "d"}, {"Е", "E"}, {"е", "e"},
		{"Ж", "J"}, {"ж", "j"}, {"З", "Z"}, {"з", "z"}, {"И", "I"}, {"и", "i"},
		{"Й", "Y"}, {"й", "y"}, {"К", "K"}, {"к", "k"}, {"Л", "L"}, {"л", "l"},
		{"М", "M"}, {"м", "m"}, {"Н", "N"}, {"н", "n"}, {"О", "O"}, {"о", "o"},
		{"П", "P"}, {"п", "p"}, {"Р", "R"}, {"р", "r"}, {"С", "S"}, {"с", "s"},
		{"Т", "T"}, {"т", "t"}, {"У", "U"}, {"у", "u"}, {"Ф", "F"}, {"ф", "f"},
		{"Х", "X"}, {"х", "x"}, {"Э", "E"}, {"э", "e"},
	}

	result := text

	// Avval 2 xonali va maxsus belgilarni almashtiramiz
	for _, r := range specificReplacements {
		result = strings.ReplaceAll(result, r.cyr, r.lat)
	}

	// Keyin oddiy 1 xonali harflarni almashtiramiz
	for _, r := range baseReplacements {
		result = strings.ReplaceAll(result, r.cyr, r.lat)
	}

	return result
}

func LatinToCyrillic(text string) string {
	replacements := []struct{ lat, cyr string }{
		{"Yo", "Ё"}, {"yo", "ё"},
		{"Yu", "Ю"}, {"yu", "ю"},
		{"Ya", "Я"}, {"ya", "я"},
		{"Ch", "Ч"}, {"ch", "ч"},
		{"Sh", "Ш"}, {"sh", "ш"},
		{"Oʻ", "Ў"}, {"oʻ", "ў"},
		{"O'", "Ў"}, {"o'", "ў"},
		{"Gʻ", "Ғ"}, {"gʻ", "ғ"},
		{"G'", "Ғ"}, {"g'", "ғ"},
		{"Ts", "Ц"}, {"ts", "ц"},
		{"Ç", "Ч"}, {"ç", "ч"},
		{"Ş", "Ш"}, {"ş", "ш"},
		{"Ö", "Ў"}, {"ö", "ў"},
		{"Ğ", "Ғ"}, {"ğ", "ғ"},
		{"A", "А"}, {"a", "а"}, {"B", "Б"}, {"b", "б"}, {"V", "В"}, {"v", "в"},
		{"G", "Г"}, {"g", "г"}, {"D", "Д"}, {"d", "д"}, {"E", "Е"}, {"e", "е"},
		{"J", "Ж"}, {"j", "ж"}, {"Z", "З"}, {"z", "з"}, {"I", "И"}, {"i", "и"},
		{"Y", "Й"}, {"y", "й"}, {"K", "К"}, {"k", "к"}, {"L", "Л"}, {"l", "л"},
		{"M", "М"}, {"m", "м"}, {"N", "Н"}, {"n", "н"}, {"O", "О"}, {"o", "о"},
		{"P", "П"}, {"p", "п"}, {"R", "Р"}, {"r", "р"}, {"S", "С"}, {"s", "с"},
		{"T", "Т"}, {"t", "т"}, {"U", "У"}, {"u", "у"}, {"F", "Ф"}, {"f", "ф"},
		{"X", "Х"}, {"x", "х"}, {"Q", "Қ"}, {"q", "қ"}, {"H", "Ҳ"}, {"h", "ҳ"},
		{"’", "Ъ"}, {"'", "Ъ"},
	}

	result := text
	for _, r := range replacements {
		result = strings.ReplaceAll(result, r.lat, r.cyr)
	}
	return result
}