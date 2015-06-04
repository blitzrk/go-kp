package dp

import "testing"

var tests = []struct {
	in  *Parameters
	out int
}{
	{
		in: &Parameters{
			Scores:    []int{3, 4, 6},
			Weights:   []int{3, 4, 6},
			MaxWeight: 7,
		},
		out: 7,
	},
	{
		in: &Parameters{
			Scores:    Round([]float64{0.3, 0.4, 0.6}, 1),
			Weights:   []int{3, 4, 6},
			MaxWeight: 7,
		},
		out: 7,
	},
	{
		in: &Parameters{
			Scores:    []int{3, 4, 5, 6},
			Weights:   []int{2, 3, 4, 5},
			MaxWeight: 5,
		},
		out: 7,
	},
	{
		in: &Parameters{
			Scores:    []int{1, 6, 18, 22, 28},
			Weights:   []int{1, 2, 5, 6, 7},
			MaxWeight: 11,
		},
		out: 40,
	},
}

func TestMax(t *testing.T) {
	for _, v := range tests {
		outVal, outItems := Max(v.in)
		t.Logf("Val: %v, Items: %v\n", outVal, outItems)
		if outVal != v.out {
			t.Errorf("Got %v, expected %v\n", outVal, v.out)
		}
	}
}
