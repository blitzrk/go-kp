// Implements a general genetic algorithm for use in solving 0-1 knapsack
// problems.
package ga

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
)

// A chromosome is represented as a byte slice where each byte stores a
// character from the genome's alphabet. So for a binary problem, it may store
// 0x0s and 0x1s, and for a genetics problem it may store A, T, C, and G (as
// 0x0, 0x1, 0x2, 0x3).
type Chromosome []byte

// This interface is for a construct that can find the fitness value of a given
// chromosome. The reason for having a struct to perform the calculations
// instead of a closure (because both can manage caching) is to allow wrapping
// in related functions that use Fitness such as greedy algorithms for
// optimizing initial populations.
type Performance interface {
	Fitness(Chromosome) float64
}

func (c Chromosome) String() string {
	var b bytes.Buffer
	a := []byte(c)
	f := "%#v"

	b.WriteString("[")
	b.WriteString(fmt.Sprintf(f, a[0]))
	for i := 1; i < len(a); i++ {
		b.WriteString(fmt.Sprintf(", "+f, a[i]))
	}
	b.WriteString("]")
	return b.String()
}

// All available parameters for tweaking the way that the algorithm works are
// manipulated through Parameters. Uint varieties are used to help inform what
// the expected size of each parameter should be.
type Parameters struct {
	Performance
	ChromLen   uint
	ChromChars uint8
	Pop        uint
	Elite      uint8
	Crosses    uint8
	CrossProb  float64
	MutateProb float64
	MaxGens    uint32
}

func randChromosomeFunc(chars, clen int, p Performance) func() Chromosome {
	return func() Chromosome {
		c := make(Chromosome, clen)
		for p.Fitness(c) == 0 {
			for i := 0; i < len(c); i++ {
				c[i] = byte(rand.Intn(chars))
			}
		}
		return c
	}
}

// Each run returns a single RunResult, which contains the best value obtained
// and the corresponding Chromosome.
type RunResult struct {
	FitVal     float64
	Chromosome Chromosome
}

type byFitness []*RunResult

func (r byFitness) Len() int           { return len(r) }
func (r byFitness) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r byFitness) Less(i, j int) bool { return r[i].FitVal < r[j].FitVal }

// Sorts results by best fitness value
func SortResults(rs []*RunResult) {
	sort.Sort(sort.Reverse(byFitness(rs)))
}

// This is the main function for running the algorithm.
func Run(p *Parameters) (*RunResult, error) {
	for gen := 0; gen < int(p.MaxGens); gen++ {
		// Generate initial population
		cf := randChromosomeFunc(int(p.ChromChars), int(p.ChromLen), p.Performance)
		g := NewInitGen(int(p.Pop), cf)
		if gp, ok := interface{}(p).(GreedyPerformance); ok {
			ImproveInitGen(g, gp)
		}

		// Select portion of population to breed
		breeders := g.Select(int(p.Elite), p.Performance)

		// TODO(ben): Breed
		_ = breeders
	}
	return nil, nil
}
