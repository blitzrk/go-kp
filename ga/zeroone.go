package ga

import "sort"

// Conforms to the GreedyPerformance
type ZeroOneFit struct {
	cache map[string]float64

	Scores  []float64
	Weights []float64
	MaxW    float64
}

func (b *ZeroOneFit) Fitness(cr Chromosome) float64 {
	if b.cache == nil {
		b.cache = make(map[string]float64)
	}
	if v, ok := b.cache[string(cr)]; ok {
		return v
	}

	var sumS float64
	var sumW float64
	for i, v := range cr {
		if v == 0x1 {
			sumS += b.Scores[i]
			sumW += b.Weights[i]
		}
	}

	// Zero fitness if infeasible
	if sumW > b.MaxW {
		sumS = 0
	}

	b.cache[string(cr)] = sumS
	return sumS
}

func (b *ZeroOneFit) Greedy() Chromosome {
	pairs := make([]pair, len(b.Scores))
	for i, v := range b.Scores {
		pairs[i] = pair{i, v}
	}
	sort.Sort(sort.Reverse(byScore(pairs)))

	var total float64
	var curr int
	best := make(Chromosome, len(pairs))
	for i := 0; total <= b.MaxW; i++ {
		curr = pairs[i].item
		best[curr] = 0x1
		total += b.Weights[curr]
	}
	best[curr] = 0x0
	return best
}
