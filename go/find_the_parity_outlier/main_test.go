package main

import "testing"

type Entry struct {
	input  []int
	result int
}

func Test_findOutlier(t *testing.T) {
	entries := []Entry{
		{input: []int{2, 4, 0, 100, 4, 11, 2602, 36}, result: 11},
		{input: []int{160, 3, 1719, 19, 11, 13, -21}, result: 160},
	}

	t.Run("should find ouliers", func(t *testing.T) {
		for _, e := range entries {
			result := FindOutlier(e.input)
			if result != e.result {
				t.Fatalf("failed to identify outlier at entry %+v, reseived %+v instead of %+v", e.input[:4], result, e.result)
			}
		}
	})
}
