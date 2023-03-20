package main

const (
	even int = iota
	uneven
	not_set
)

func isEven(i int) bool {
	return i%2 == 0
}

func FindOutlier(integers []int) int {
	i := len(integers) - 1

	if len(integers) < 3 {
		return 0
	}
	trend := not_set

	isEven0 := isEven(integers[i])
	isEven1 := isEven(integers[i-1])

	if isEven0 == isEven1 {
		if isEven0 == false {
			trend = uneven
		} else {
			trend = even
		}
	} else {
		isEven2 := isEven(integers[i-2])
		if isEven2 == isEven0 {
			return integers[i-1]
		}
		if isEven2 == isEven1 {
			return integers[i]
		}
	}

	i -= 2
	for i >= 0 {
		switch trend {
		case even:
			if !isEven(integers[i]) {
				return integers[i]
			}
		case uneven:
			if isEven(integers[i]) {
				return integers[i]
			}
		}
		i--
	}

	return 0
}

func main() {
}
