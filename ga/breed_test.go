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
			t.Errorf("Failed test #%v: Expected %v, got %v.\n", num+1, test.out, test.in)
		}
	}
}

func TestPairOff(t *testing.T) {
	tests := []struct {
		in  *Breeder
		out []*breedPair
	}{
		{
			&Breeder{
				nil,
				&Generation{
					generation{
						&TestCM{0x1, 0x0, 0x0},
						&TestCM{0x0, 0x1, 0x0},
						&TestCM{0x0, 0x0, 0x1},
					},
					nil,
				},
			},
			[]*breedPair{
				&breedPair{
					&TestCM{0x1, 0x0, 0x0},
					&TestCM{0x0, 0x1, 0x0},
				},
				&breedPair{
					&TestCM{0x0, 0x1, 0x0},
					&TestCM{0x0, 0x0, 0x1},
				},
			},
		},
		{
			&Breeder{
				nil,
				&Generation{
					generation{
						&TestCM{0x1, 0x0, 0x0},
						&TestCM{0x0, 0x1, 0x0},
					},
					nil,
				},
			},
			[]*breedPair{
				&breedPair{
					&TestCM{0x1, 0x0, 0x0},
					&TestCM{0x0, 0x1, 0x0},
				},
			},
		},
	}

	for num, test := range tests {
		out := test.in.pairOff()
		if !breedPairsEqual(out, test.out) {
			t.Errorf("Failed test #%v: Expected %v, got %v.\n", num+1, test.out, out)
		}
	}
}

func TestBreed(t *testing.T) {
	tests := []struct {
		in        *Breeder
		inNCross  int
		inPCross  float64
		inPMutate float64
		out       generation
	}{
		{
			&Breeder{
				rand.New(&TestRandSource{}),
				&Generation{
					generation{
						&TestCM{0x1, 0x0, 0x0, 0x0},
						&TestCM{0x0, 0x1, 0x0, 0x0},
					}, nil,
				},
			},
			0,
			0,
			0,
			generation{
				&TestCM{0x1, 0x0, 0x0, 0x0},
				&TestCM{0x0, 0x1, 0x0, 0x0},
			},
		},
	}

	for num, test := range tests {
		out := test.in.Breed(test.inNCross, test.inPCross, test.inPMutate)
		if !generationsEqual(out.gen, test.out) {
			t.Errorf("Failed test #%v: Expected %v, got %v.\n", num+1, test.out, out)
		}
	}
}
