package ga

import (
	"math/rand"
	"testing"
)

type TestCMMut struct {
	*TestCM
}

func (cm TestCMMut) MutateChar(i int) {
	(*cm.TestCM)[i] = 1 - (*cm.TestCM)[i]
}

func (cm1 TestCMMut) Cross(locus int, cm2 ChromosomeModel) (ChromosomeModel, ChromosomeModel) {
	return nil, nil
}

// Source that alternates between 0 and 0.75
type TestRandSource struct {
	last int64
}

func (s *TestRandSource) Seed(seed int64) {}
func (s *TestRandSource) Int63() int64 {
	if s.last == 0 {
		s.last = 1<<62 + 1<<61
		return 1<<62 + 1<<61
	} else {
		s.last = 0
		return 0
	}
}

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
