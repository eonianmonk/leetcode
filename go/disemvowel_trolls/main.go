package main

import "strings"

var (
	vovels    = "aeiou"
	vovelsCap = strings.ToUpper(vovels)
)

func isVovel(r rune) bool {
	if r >= 'a' && r < 'z' {
		return strings.ContainsRune(vovels, r)
	}
	if r > 'A' && r < 'Z' {
		return strings.ContainsRune(vovelsCap, r)
	}
	return false
}

func Disemvowel(source string) string {
	sb := strings.Builder{}
	for _, r := range source {
		if !isVovel(r) {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func main() {

}
