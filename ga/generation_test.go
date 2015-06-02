package ga

import "testing"

func TestNewInitGen(t *testing.T) {
	tests := []struct {
		pop int
		fun func() Chromosome
	}{
		{0, randChromosomeFunc(2, 7)},
		{5, randChromosomeFunc(1, 5)},
		{9, randChromosomeFunc(255, 0)},
	}

	for _, test := range tests {
		if NewInitGen(test.pop, test.fun) == nil {
			t.Error("NewIntGen failed")
		}
	}
}
