package ga

import (
	"sort"
	"testing"
)

var tests = []struct {
	in  metadata
	out metadata
}{
	{
		metadata{
			data{1, 1.5},
			data{2, 2.0},
		},
		metadata{
			data{2, 2.0},
			data{1, 1.5},
		},
	},
	{
		metadata{
			data{1, 6.6},
			data{2, 12.0},
			data{3, 7.0},
			data{4, 5.5},
			data{5, 8.0},
		},
		metadata{
			data{2, 12.0},
			data{5, 8.0},
			data{3, 7.0},
			data{1, 6.6},
			data{4, 5.5},
		},
	},
}

func metadataEqual(ps1, ps2 metadata) bool {
	if len(ps1) != len(ps2) {
		return false
	}

	for i := 0; i < len(ps1); i++ {
		if ps1[i].item != ps2[i].item || ps1[i].score != ps2[i].score {
			return false
		}
	}
	return true
}

func TestSort(t *testing.T) {
	for i := 0; i < len(tests); i++ {
		sort.Sort(sort.Reverse(byScore(tests[i].in)))
		if !metadataEqual(tests[i].in, tests[i].out) {
			t.Errorf("Sorting failed, got %v, expected %v", tests[i].in, tests[i].out)
		}
	}
}

func TestMergeSortedDesc(t *testing.T) {
	for i := 1; i < len(tests); i++ {
		expect := make(metadata, len(tests[i].out))
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
