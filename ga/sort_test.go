package ga

import (
	"sort"
	"testing"
)

var tests = []struct {
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

func (ps1 pairs) equals(ps2 pairs) bool {
	if len(ps1) != len(ps2) {
		return false
	}

	for i := 0; i < len(ps1); i++ {
		if ps1[i].item != ps2[i].item {
			return false
		}
	}
	return true
}

func TestSort(t *testing.T) {
	for i := 0; i < len(tests); i++ {
		sort.Sort(sort.Reverse(byScore(tests[i].in)))
		if !tests[i].in.equals(tests[i].out) {
			t.Errorf("Sorting failed, got %v, expected %v", tests[i].in, tests[i].out)
		}
	}
}

func TestMergeSortedDesc(t *testing.T) {
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
