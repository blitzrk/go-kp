// Implements a general genetic algorithm for use in solving 0-1 knapsack
// problems.
package ga

import (
	"bytes"
	"fmt"
	"sort"
)

// A chromosome is represented as a byte slice where each byte stores a
// character from the genome's alphabet. So for a binary problem, it may store
// 0x0s and 0x1s, and for a genetics problem it may store A, T, C, and G (as
// 0x0, 0x1, 0x2, 0x3).
type Chromosome []byte

// Types that supertype (or embed) Chromosome should implement the
// ChromosomeModel interface.
type ChromosomeModel interface {
	Key() string
	Len() int
	Loc(int) byte
	MutateChar(int)
	Cross(int, ChromosomeModel) (ChromosomeModel, ChromosomeModel)
}

// This interface is for a construct that can find the fitness value of a given
// chromosome. The reason for having a struct to perform the calculations
// instead of a closure (because both can manage caching) is to allow wrapping
// in related functions that use Fitness such as greedy algorithms for
// optimizing initial populations. It is also a generator for random
// Chromosomes. Generation is tied to Fitness, because Rand should not return
// a non-viable Chromosome.
type Performance interface {
	Fitness(ChromosomeModel) float64
	Rand() ChromosomeModel
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
	Perf       Performance // Struct with fitness information
	Loci       uint        // Length of chromosome in characters
	InitPop    uint        // How many (random) chromosomes to start with
	Elite      uint8       // How many of the best chromosomes to keep from each generation
	Crosses    uint8       // Max number of locations to cross DNA during breeding
	CrossProb  float64     // Probability of a single cross happening
	MutateProb float64     // Probability of a mutation happening at a locus
	MaxGens    uint32      // Stop the algorithm after this many gens, even w/o convergence
}

// Each run returns a single RunResult, which contains the best value obtained
// and the corresponding Chromosome.
type RunResult struct {
	FitVal     float64
	Chromosome ChromosomeModel
}

func (rr *RunResult) String() string {
	return fmt.Sprintf("Chromosome %v, score %v", rr.Chromosome, rr.FitVal)
}

type byFitness []*RunResult

func (r byFitness) Len() int           { return len(r) }
func (r byFitness) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r byFitness) Less(i, j int) bool { return r[i].FitVal < r[j].FitVal }

// Sorts results by best fitness value
func SortResults(rs []*RunResult) {
	sort.Sort(sort.Reverse(byFitness(rs)))
}

func shouldStop(g *Generation, p Performance) bool {
	// TODO(ben): population becomes significantly similar
	return false
}

func reportBest(g *Generation, p Performance) *RunResult {
	if g.gen == nil || len(g.gen) == 0 {
		return nil
	}

	r := g.rank(p)
	best := g.gen[r[0].item]
	score := p.Fitness(best)

	return &RunResult{
		score,
		best,
	}
}

// This is the main function for running the algorithm.
func Run(p *Parameters) (*RunResult, error) {
	// Generate initial population
	g := NewInitGen(int(p.InitPop), p.Perf)
	if gp, ok := interface{}(p.Perf).(GreedyPerformance); ok {
		ImproveInitGen(g, gp)
	}

	curr := g
	for gen := 0; gen < int(p.MaxGens); gen++ {
		// Check stopping conditions
		if shouldStop(curr, p.Perf) {
			break
		}

		// Select portion of population to breed
		parents := curr.Select(int(p.Elite), p.Perf)

		// Breed
		children := parents.Breed(int(p.Crosses), p.CrossProb, p.MutateProb)
		curr = children
	}

	return reportBest(curr, p.Perf), nil
}
