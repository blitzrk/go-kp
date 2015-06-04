package ga

import "testing"

var perf = &ZeroOneFit{
	Scores:  []float64{3, 4, 6},
	Weights: []float64{3, 4, 6},
	MaxW:    7,
}

func TestNewInitGen(t *testing.T) {
	tests := []struct {
		pop int
		fun func() Chromosome
	}{
		{0, randChromosomeFunc(2, 3, perf)},
		{5, randChromosomeFunc(2, 3, perf)},
	}

	for _, test := range tests {
		if NewInitGen(test.pop, test.fun) == nil {
			t.Error("NewIntGen failed")
		}
	}
}
