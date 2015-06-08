package zeroone

import (
	"testing"

	"github.com/blitzrk/go-kp/ga"
)

func chromosomesEqual(cm1, cm2 ga.ChromosomeModel) bool {
	if cm1.Len() != cm2.Len() {
		return false
	}
	for i := 0; i < cm1.Len(); i++ {
		if cm1.Loc(i) != cm2.Loc(i) {
			return false
		}
	}
	return true
}

func TestGreedy(t *testing.T) {
	tests := []struct {
		in  ga.GreedyPerformance
		out ga.ChromosomeModel
	}{
		{
			&Fit{
				Scores:  []float64{3, 4, 6},
				Weights: []float64{3, 4, 6},
				MaxW:    7,
			},
			&Chromosome{0x0, 0x0, 0x1},
		},
	}

	for num, test := range tests {
		out := test.in.Greedy()
		if !chromosomesEqual(out, test.out) {
			t.Errorf("Test #%v failed: Expected %v, got %v.\n", num+1, test.out, out)
		}
	}
}
