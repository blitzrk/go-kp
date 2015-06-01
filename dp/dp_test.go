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
}

func TestMax(t *testing.T) {
	for _, v := range tests {
		out := Max(v.in)
		if out != v.out {
			t.Errorf("Got %v, expected %v", out, v.out)
		}
	}
}
