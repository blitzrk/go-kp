package ga

import (
	"testing"
)

var testParam = &Parameters{
	ChromLen:   3,
	ChromChars: 2,
	Pop:        3,
	Elite:      2,
	Crosses:    1,
	CrossProb:  0.85,
	MutateProb: 0.001,
	MaxGens:    5,
}

var testGen = generation{
	Chromosome{0x1, 0x1, 0x0, 0x1},
	Chromosome{0x0, 0x1, 0x1, 0x0},
	Chromosome{0x1, 0x0, 0x0, 0x0},
	Chromosome{0x0, 0x0, 0x0, 0x1},
}

// (cd ga; go test -v)
func TestNewInitGen(t *testing.T) {
	t.Log(NewInitGen(testParam))
}
