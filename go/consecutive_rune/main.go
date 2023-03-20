package main

type Result struct {
	C rune
	L int // count
}

func toLowerCase(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r
	}
	if r >= 'A' && r <= 'Z' {
		const interval rune = 'A' - 'a'
		return r - interval
	}
	return r
}

func LongestRepetition(text string) Result {
	print(text)
	if len(text) == 0 {
		return Result{}
	}

	var maxRune rune
	var maxCount int = 0

	var consRune rune = 0
	var consCount int = 0
	for _, r := range text {
		r = toLowerCase(r)
		switch {
		case r == consRune:
			consCount++
		case r != consRune:
			if consCount > maxCount {
				maxCount = consCount
				maxRune = consRune
			}
			consRune = r
			consCount = 1
		}
	}
	if consCount > maxCount {
		maxRune = consRune
		maxCount = consCount
	}
	return Result{C: maxRune, L: maxCount}
}

func main() {
	print(len(""))
}
