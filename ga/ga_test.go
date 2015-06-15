package ga

import (
	"testing"

	"github.com/blitzrk/go-kp/ga/zo"
)

var testParam = &Parameters{
	Perf: &zo.Perf{
		Scores:  []float64{1, 1, 1, 1},
		Weights: []float64{1, 1, 1, 1},
		MaxW:    4,
	},
	InitPop:    8,
	Elite:      2,
	Crosses:    1,
	CrossProb:  0.85,
	MutateProb: 0.001,
	MaxGens:    10,
}

func TestRun(t *testing.T) {
	if rr, _ := Run(testParam); rr.FitVal != 4 {
		t.Errorf("Did not find optimal soln 4, got: %v", rr)
	}
}
