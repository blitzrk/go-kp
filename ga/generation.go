package ga

import (
	"fmt"
	"log"
	"sort"
)

type generation []Chromosome

// A Generation is a collection of chromosomes with a method of selecting,
// based on fitness, which chromosomes should be used as parents for the
// breeding phase.
type Generation struct {
	generation
	meta metadata
}

// Returns a string representing the array of chromosomes.
func (g *Generation) String() string {
	return fmt.Sprint(g.generation)
}

// Returns information on the ordering of the Generation's chromosomes by
// fitness.
func (g *Generation) rank(p Performance) metadata {
	if g.meta != nil {
		return g.meta
	}

	ps := make(metadata, len(g.generation))
	for i, c := range g.generation {
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
		pick[i] = g.generation[v]
	}
	return pick
}

// Separates the n chromosomes with highest fitness and the rest as two
// separate Generations and returns pointers to them.
func (g *Generation) extractElite(n int, p Performance) (*Generation, *Generation) {
	if n > len(g.generation) {
		log.Printf(`Warning: Elitism set to %v, but generation only has %v chromosomes.
			Therefore, the entire generation will be selected as parents.\n`, n,
			len(g.generation))
		n = len(g.generation)
	}

	elite := make(generation, n)
	rem := make(generation, len(g.generation)-n)

	ranked := g.rank(p)
	sortedChroms := g.cherryPick(ranked.Items())
	copy(elite, sortedChroms[:n])
	copy(rem, sortedChroms[n:])

	return &Generation{elite, ranked.Subset(0, n)},
		&Generation{rem, ranked.Subset(n, len(ranked))}
}

// Implements the roulette-wheel selection method for deciding
// probabilistically which chromosomes to breed. Alternate selection methods
// exist, but for now roulette is built in and is the only option.
func (g *Generation) roulette(p Performance) *Generation {
	// TODO(ben): implement roulette method

	return g
}

// Merges two generations and returns the new one instance pointer. The intent
// of this method is to be used to recombine elite chromosomes with the other
// selected chromosomes to create a selected parent Generation. Because each
// generation should be ranked already, and the merged ranks will be needed,
// an optimization is applied for determining the new ranks faster.
func (g1 *Generation) merge(g2 *Generation) *Generation {
	m := make(generation, len(g1.generation)+len(g2.generation))
	copy(m[:len(g1.generation)], g1.generation)
	copy(m[len(g1.generation):], g2.generation)

	if g1.meta == nil || g2.meta == nil {
		return &Generation{m, nil}
	}
	return &Generation{m, g1.meta.MergeSortedDesc(g2.meta)}
}

// Selects the n most fit chromosomes and some of the remaining (using the
// roulette method) based on fitness to be used as the parent generation.
func (g *Generation) Select(nElite int, p Performance) *Generation {
	elite, rem := g.extractElite(nElite, p)
	sel := rem.roulette(p)
	return sel.merge(elite)
}

// Creates a random generation of parent candidates of popSize using the given
// function to generate a random chromosome.
func NewInitGen(popSize int, randChrom func() Chromosome) *Generation {
	gen := make(generation, popSize)

	for i := 0; i < len(gen); i++ {
		gen[i] = randChrom()
	}

	return &Generation{gen, nil}
}

// If a greedy algorithm for finding the most fit chromosome (utilizing the
// Fitness method) is embedded with the Fitness method in the Parameters, it
// can be used to potentially improve the randomly generated initial
// generation.
type GreedyPerformance interface {
	Fitness(Chromosome) float64
	Greedy() Chromosome
}

// Tries to improve the initial generation by using a (hopefully fast) greedy
// algorithm to add a (hopefully good) chromosome and then removing the worst.
func ImproveInitGen(gen *Generation, gp GreedyPerformance) {
	g := &Generation{generation{gp.Greedy()}, nil}
	gen = gen.merge(g)
	gen.rank(gp)

	// Known memory leak: but only 2 meta and only run once
	worst := gen.meta[len(gen.meta)-1].item
	gen = &Generation{
		append(gen.generation[:worst], gen.generation[worst+1:]...),
		gen.meta[:len(gen.meta)-1],
	}
}
