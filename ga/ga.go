package ga

import ()

type Chromosome []byte

type Generation []Chromosome

type Parameters struct {
	Pop        uint64
	Elite      uint8
	Crosses    uint8
	CrossProb  float64
	MutateProb float64
	Scores     []float64
	Weights    []float64
	MaxWeight  float64
	MaxGens    uint32
}

func (p *Parameters) Fitness(cr Chromosome) float64 {
	var sum float64
	for i, v := range cr {
		if v == 0x1 {
			sum += p.Scores[i]
		}
	}
	return sum
}

type pair struct {
	item  int
	score float64
}

type byScore []pair

func (s byScore) Len() int           { return len(s) }
func (s byScore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byScore) Less(i, j int) bool { return s[i].score < s[j].score }

func (p *Parameters) findGreedyMax() Chromosome {
	pairs := make([]pair, len(p.Scores))
	for i, v := range p.Scores {
		pairs[i] = pair{i, v}
	}
	sort.Sort(byScore(pairs))

	var total float64
	best := make(Chromosome, len(p.Scores))
	for i := 0; total <= p.MaxWeight; i++ {
		best[pairs[i].item] = 0x1
		total += p.Weights[pairs[i].item]
	}
	return best[:len(best)-1]
}

func NewRandPop() {

}
