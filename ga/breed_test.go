package ga

import (
	"math/rand"
	"testing"
)

func TestMutateChromosome(t *testing.T) {
	tests := []struct {
		in   ChromosomeModel
		rand *rand.Rand
		prob float64
		out  ChromosomeModel
	}{
		{
			TestCMMut{&TestCM{0x0, 0x0, 0x0, 0x0}},
			rand.New(&TestRandSource{}),
			0.5,
			TestCMMut{&TestCM{0x0, 0x1, 0x0, 0x1}},
		},
	}

	for num, test := range tests {
		mutateChromosome(test.in, test.prob, test.rand)
		if !chromosomesEqual(test.in, test.out) {
			t.Errorf("Failed test #%v: Expected %v, got %v.\n", num, test.out, test.in)
		}
	}
}

func TestBreed(t *testing.T) {
	t.Log("TODO")
}
