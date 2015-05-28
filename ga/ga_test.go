package ga

import (
	"testing"
)

var testP = &Parameters{
	ChromLen:   3,
	Pop:        3,
	Elite:      2,
	Crosses:    1,
	CrossProb:  0.85,
	MutateProb: 0.001,
	MaxGens:    5,
}

func TestPop(t *testing.T) {
	t.Log(NewRandPop(testP))
}
