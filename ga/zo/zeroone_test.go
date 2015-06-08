package zeroone

import (
	"testing"
)

var testZeroOneP = &Parameters{
	Pop:        30,
	Elite:      2,
	Crosses:    1,
	CrossProb:  0.85,
	MutateProb: 0.001,
	MaxGens:    5,
	Performance: &Fit{
		Scores:  []float64{3, 4, 6},
		Weights: []float64{3, 4, 6},
		MaxW:    7,
	},
}

var testGen = &Generation{generation{
	Chromosome{0x1, 0x1, 0x0, 0x1},
	Chromosome{0x0, 0x1, 0x1, 0x0},
	Chromosome{0x1, 0x0, 0x0, 0x0},
	Chromosome{0x0, 0x0, 0x0, 0x1},
}, nil}

func TestGreedy(t *testing.T) {
	if gp, ok := interface{}(testZeroOneP.Performance).(GreedyPerformance); ok {
		got := gp.Greedy()
		expect := Chromosome([]byte{0, 0, 1})

		if string(got) != string(expect) {
			t.Errorf("Got %v, expected %v", got, expect)
		}
	} else {
		t.Error("ZeroOneFit should implement GreedyPerformance")
	}
}
