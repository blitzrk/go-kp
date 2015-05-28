package ga

import "sort"

type BinaryFit struct {
	cache map[string]float64

	Scores  []float64
	Weights []float64
	MaxW    float64
}

func (b *BinaryFit) Fitness(cr Chromosome) float64 {
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

type pair struct {
	item  int
	score float64
}

type byScore []pair

func (s byScore) Len() int           { return len(s) }
func (s byScore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byScore) Less(i, j int) bool { return s[i].score < s[j].score }

func (b *BinaryFit) Greedy() Chromosome {
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
