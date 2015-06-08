package ga

import (
	"sort"
	"testing"
)

func metadataEqual(ps1, ps2 metadata) bool {
	if ps1 == nil && ps2 == nil {
		return true
	}
	if ps1 == nil || ps2 == nil {
		return false
	}
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
	tests := []struct {
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

	for num, test := range tests {
		out := make(metadata, len(test.in))
		copy(out, test.in)
		sort.Sort(sort.Reverse(byScore(out)))
		if !metadataEqual(out, test.out) {
			t.Errorf("Sort #%v failed, got %v, expected %v", num+1, out, test.out)
		}
	}
}
