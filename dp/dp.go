// Implements a dynamic programming algorithm for solving any 0-1 knapsack
// problem. It grows in space and time with O(n*W).
package dp

import (
	"math"
	"sort"
)

// A helper function for converting weights and scores from []float64 to []int.
// For example, Round(dollars, 2) incorporates cents and Round(dollars, 0)
// ignores (truncates) change. This may inherently alter the units of the
// solution's max value.
func Round(floats []float64, n int) (ints []int) {
	ints = make([]int, len(floats))
	for i := 0; i < len(ints); i++ {
		r := math.Floor(floats[i] * math.Pow(10, float64(n)))
		ints[i] = int(r)
	}

	return
}

// The necessary parameters for the DP algorithm are the scores and weights,
// aligned by index, and the maximum weight the knapsack can hold. All scores
// and weights must be integer-valued.
type Parameters struct {
	Scores    []int
	Weights   []int
	MaxWeight int
	v         [][]int
	keep      [][]int
}

func (p *Parameters) reset() {
	n := len(p.Scores)
	W := p.MaxWeight

	p.v = make([][]int, n+1)
	p.v[0] = make([]int, W+1)
	for i := 1; i <= n; i++ {
		p.v[i] = make([]int, W+1)
		for j := 0; j <= W; j++ {
			p.v[i][j] = int(math.MinInt64)
		}
	}

	p.keep = make([][]int, n)
	for i := 0; i < n; i++ {
		p.keep[i] = make([]int, W+1)
	}
}

func (p *Parameters) evalAt(i, w int) int {
	// i counts up to which item to consider 1..n
	// j counts index for weight of item 0..n-1
	j := i - 1

	prev := p.v[i-1][w]
	// check if the knapsack can even hold just this item
	if w < p.Weights[j] {
		return prev
	}

	alt := p.Scores[j] + p.v[i-1][w-p.Weights[j]]
	if alt > prev {
		p.keep[j][w] = 1
		return alt
	} else {
		return prev
	}
}

// This is the main function to perform the optimization algorithm. It returns
// the optimal value (max score total) along with an array of indices for the
// chosen items making up the optimal solution.
func Max(p *Parameters) (optVal int, soln []int) {
	p.reset()
	n := len(p.Scores)
	W := p.MaxWeight

	for i := 1; i <= n; i++ {
		for w := 0; w <= W; w++ {
			p.v[i][w] = p.evalAt(i, w)
		}
	}

	soln = make([]int, 0)
	K := W
	for i := n - 1; i >= 0; i-- {
		if p.keep[i][K] == 1 {
			soln = append(soln, i)
			K = K - p.Weights[i]
		}
	}
	sort.Ints(soln)

	optVal = p.v[n][W]
	return
}
