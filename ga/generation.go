package ga

import (
	"fmt"
	"sort"
)

type generation []ChromosomeModel

// A Generation is a collection of chromosomes with a method of selecting,
// based on fitness, which chromosomes should be used as parents for the
// breeding phase.
type Generation struct {
	gen  generation
	meta metadata
}

// Returns a string representing the array of chromosomes.
func (g *Generation) String() string {
	return fmt.Sprint(g.gen)
}

// Returns information on the ordering of the Generation's chromosomes by
// fitness.
func (g *Generation) rank(p Performance) metadata {
	if g.meta != nil {
		return g.meta
	}

	ps := make(metadata, len(g.gen))
	for i, c := range g.gen {
		ps[i] = data{i, p.Fitness(c)}
	}
	sort.Sort(sort.Reverse(byScore(ps)))
	g.meta = ps

	return ps
}

// Uses an int slice to pick chromosomes by their fitness rank in any order,
// possibly with missing or duplicate instances.
func (g *Generation) cherryPick(its []int) generation {
	pick := make(generation, len(its))
	for i, v := range its {
		pick[i] = g.gen[v]
	}
	return pick
}

// Merges two generations and returns the new one instance pointer. The intent
// of this method is to be used to recombine elite chromosomes with the other
// selected chromosomes to create a selected parent Generation. Because each
// generation should be ranked already, and the merged ranks will be needed,
// an optimization is applied for determining the new ranks faster.
func (g1 *Generation) merge(g2 *Generation) *Generation {
	m := make(generation, len(g1.gen)+len(g2.gen))
	copy(m[:len(g1.gen)], g1.gen)
	copy(m[len(g1.gen):], g2.gen)

	if g1.meta == nil || g2.meta == nil {
		return &Generation{m, nil}
	}
	return &Generation{m, g1.meta.MergeSortedDesc(g2.meta)}
}

// Creates a random generation of parent candidates of popSize using the given
// function to generate a random chromosome.
func NewInitGen(popSize, chromSize int, randChrom func(int) ChromosomeModel) *Generation {
	gen := make(generation, popSize)

	for i := 0; i < len(gen); i++ {
		gen[i] = randChrom(chromSize)
	}

	return &Generation{gen, nil}
}

// If a greedy algorithm for finding the most fit chromosome (utilizing the
// Fitness method) is embedded with the Fitness method in the Parameters, it
// can be used to potentially improve the randomly generated initial
// generation.
type GreedyPerformance interface {
	Performance
	Greedy() ChromosomeModel
}

// Tries to improve the initial generation by using a (hopefully fast) greedy
// algorithm to add a (hopefully good) chromosome and then removing the worst.
func ImproveInitGen(gen *Generation, gp GreedyPerformance) {
	g := &Generation{generation{gp.Greedy()}, nil}
	gen = gen.merge(g)

	// Known memory leak: but only 2 data structs and only run once
	gen.rank(gp)
	worst := gen.meta[len(gen.meta)-1].item
	gen = &Generation{
		append(gen.gen[:worst], gen.gen[worst+1:]...),
		gen.meta[:len(gen.meta)-1],
	}
}
