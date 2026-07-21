package translator

import "strings"


var cyrillicToLatinReplacer = strings.NewReplacer(
    "Ш", "Ş",
    "ш", "ş",
    "Ч", "Ç",
    "ч", "ç",
    "Ў", "Ö",
    "ў", "ö",
    "Ғ", "Ğ",
    "ғ", "ğ",
)


var latinToCyrillicReplacer = strings.NewReplacer(
   
    "Ş", "Ш",
    "ş", "ш",
    "Ç", "Ч",
    "ç", "ч",
    "Ö", "Ў",
    "ö", "ў",
    "Ğ", "Ғ",
    "ğ", "ғ",
    
    "Sh", "Ш",
    "sh", "ш",
    "Ch", "Ч",
    "ch", "ч",
    "O'", "Ў",
    "o'", "ў",
    "G'", "Ғ",
    "g'", "ғ",
)


func CyrillicToLatin(text string) string {
    return cyrillicToLatinReplacer.Replace(text)
}


func LatinToCyrillic(text string) string {
    return latinToCyrillicReplacer.Replace(text)
}
