package ga

import (
	"sort"
	"testing"
)

func TestMergeSortedDesc(t *testing.T) {
	tests := []struct {
		in  pairs
		out pairs
	}{
		{
			pairs{
				pair{1, 1.5},
				pair{2, 2.0},
			},
			pairs{
				pair{2, 2.0},
				pair{1, 1.5},
			},
		},
		{
			pairs{
				pair{1, 6.6},
				pair{2, 12.0},
				pair{3, 7.0},
				pair{4, 5.5},
				pair{5, 8.0},
			},
			pairs{
				pair{2, 12.0},
				pair{5, 8.0},
				pair{3, 7.0},
				pair{1, 6.6},
				pair{4, 5.5},
			},
		},
	}

	for i := 1; i < len(tests); i++ {
		expect := make(pairs, len(tests[i].out))
		copy(expect, tests[i].out)
		expect = append(expect, tests[i-1].out...)
		sort.Sort(sort.Reverse(byScore(expect)))

		got := tests[i].out.MergeSortedDesc(tests[i-1].out)
		if len(got) != (len(tests[i].out)+len(tests[i-1].out)) ||
			!sort.IsSorted(sort.Reverse(byScore(got))) {
			t.Errorf("Merge sort failed, got %v, expected %v", got, expect)
		}
	}
}
