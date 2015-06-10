package zeroone

import (
	"math/rand"
	"reflect"
	"sort"

	"github.com/blitzrk/go-kp/ga"
)

// Chromosome with two character alphabet.
type Chromosome []byte

func (c *Chromosome) Key() string      { return string([]byte(*c)) }
func (c *Chromosome) Len() int         { return len(*c) }
func (c *Chromosome) Loc(i int) byte   { return (*c)[i] }
func (c *Chromosome) MutateChar(i int) { (*c)[i] = byte(((*c)[i] + 1) % 2) }
func (c Chromosome) String() string    { return ga.Chromosome(c).String() }

func tryParseBytes(b interface{}) ([]byte, bool) {
	defer func() { recover() }()
	if bytes := reflect.ValueOf(b).Bytes(); bytes != nil {
		return bytes, true
	}
	if bytes := reflect.Indirect(reflect.ValueOf(b)).Bytes(); bytes != nil {
		return bytes, true
	}
	return nil, false
}

func (c *Chromosome) Cross(loc int, cm ga.ChromosomeModel) (ga.ChromosomeModel, ga.ChromosomeModel) {
	var child1, child2 Chromosome
	c1 := *c
	var c2 []byte
	if bytes, ok := tryParseBytes(cm); ok {
		c2 = bytes
	} else {
		c2 = make([]byte, cm.Len())
		for i, _ := range c2 {
			c2[i] = cm.Loc(i)
		}
	}

	copy(child1[:loc], c1[:loc])
	copy(child2[:loc], c2[:loc])
	copy(child1[loc:], c2[loc:])
	copy(child2[loc:], c1[loc:])

	return &child1, &child2
}

// Creates a function that produces a random Chromosome
func RandFactory(clen uint, p ga.Performance) func() ga.ChromosomeModel {
	return func() ga.ChromosomeModel {
		c := make(Chromosome, clen)
		for p.Fitness(&c) == 0 {
			for i := 0; i < len(c); i++ {
				c[i] = byte(rand.Intn(2))
			}
		}
		return &c
	}
}

// Conforms to the GreedyPerformance interface
type Perf struct {
	cache map[string]float64

	Scores  []float64
	Weights []float64
	MaxW    float64
}

func (perf *Perf) Fitness(cm ga.ChromosomeModel) float64 {
	if perf.cache == nil {
		perf.cache = make(map[string]float64)
	}
	if v, ok := perf.cache[cm.Key()]; ok {
		return v
	}

	var sumS float64
	var sumW float64
	for i := 0; i < cm.Len(); i++ {
		if cm.Loc(i) == 0x1 {
			sumS += perf.Scores[i]
			sumW += perf.Weights[i]
		}
	}

	// Zero fitness if infeasible
	if sumW > perf.MaxW {
		sumS = 0
	}

	perf.cache[cm.Key()] = sumS
	return sumS
}

func (perf *Perf) Rand() ga.ChromosomeModel {
	clen := len(perf.Scores)
	c := make(Chromosome, clen)
	for perf.Fitness(&c) == 0 {
		for i := 0; i < len(c); i++ {
			c[i] = byte(rand.Intn(2))
		}
	}
	return &c
}

type data struct {
	item  int
	score float64
}

type byScore []data

func (s byScore) Len() int           { return len(s) }
func (s byScore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byScore) Less(i, j int) bool { return s[i].score < s[j].score }

func (perf *Perf) Greedy() ga.ChromosomeModel {
	meta := make([]data, len(perf.Scores))
	for i, v := range perf.Scores {
		meta[i] = data{i, v}
	}
	sort.Sort(sort.Reverse(byScore(meta)))

	var total float64
	var curr int
	best := make([]byte, len(meta))
	for i := 0; total <= perf.MaxW; i++ {
		curr = meta[i].item
		best[curr] = 0x1
		total += perf.Weights[curr]
	}
	best[curr] = 0x0

	res := Chromosome(best)
	return &res
}
