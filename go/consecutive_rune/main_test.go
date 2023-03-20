package main

import "testing"

type Entry struct {
	Input        string
	WantedResult Result
}

func Test_countConsecutives(t *testing.T) {

	t.Run("count consecutives", func(t *testing.T) {
		entries := []Entry{
			{Input: "bbbaaabaaaa", WantedResult: Result{'a', 4}},
			{Input: "cbdeuuu900", WantedResult: Result{'u', 3}},
		}

		for _, e := range entries {
			received := LongestRepetition(e.Input)
			if received != e.WantedResult {
				t.Fatalf("failed to calculate repetition in %s: got %+v, wanted %+v", e.Input, received, e.WantedResult)
			}
		}
	})
}
