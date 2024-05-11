package word

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func UnderScoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	titleCaser := cases.Title(language.English)
	s = titleCaser.String(s)
	return strings.Replace(s, " ", "", -1)
}

func UnderScoreToLowerCamelCase(s string) string {
	s = UnderScoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}


func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return output 
}
