package dp

import "math"

type Parameters struct {
	Scores    []int
	Weights   []int
	MaxWeight int
	v         [][]int
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
}

func (p *Parameters) evalAt(i, w int) int {
	prev := p.v[i-1][w]
	if w < p.Weights[i-1] {
		return prev
	}
	alt := p.Scores[i-1] + p.v[i-1][w-p.Weights[i-1]]

	if alt > prev {
		return alt
	} else {
		return prev
	}
}

func Max(p *Parameters) int {
	p.reset()

	var best int
	for i := 1; i <= len(p.Scores); i++ {
		for w := 0; w <= p.MaxWeight; w++ {
			p.v[i][w] = p.evalAt(i, w)
			if p.v[i][w] > best {
				best = p.v[i][w]
			}
		}
	}

	return best
}
