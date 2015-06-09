package ga

import (
	"log"
	"math/rand"
)

// Selects the n most fit chromosomes and some of the remaining (using the
// roulette method) based on fitness to be used as the parent generation.
func (g *Generation) Select(nElite int, p Performance) *Breeder {
	elite, rem := g.extractElite(nElite, p)
	sel := rem.roulette()
	return NewBreeder(elite.append(sel))
}

// Separates the n chromosomes with highest fitness and the rest as two
// separate Generations and returns pointers to them.
func (g *Generation) extractElite(n int, p Performance) (*Generation, *Generation) {
	if n > len(g.gen) {
		log.Printf(`Warning: Elitism set to %v, but generation only has %v chromosomes.
			Therefore, the entire generation will be selected as parents.\n`, n,
			len(g.gen))
		n = len(g.gen)
	}

	elite := make(generation, n)
	rem := make(generation, len(g.gen)-n)

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
func (g *Generation) roulette() *Generation {
	// Determine cutoffs for CDF
	scores := g.meta.Scores()
	var total float64
	for _, v := range scores {
		total += v
	}
	cutoffs := make([]float64, len(g.meta))
	for i, _ := range cutoffs {
		cutoffs[i] = scores[i] / total
	}

	// Make n spins to select up to n chromosomes
	n := len(g.gen)
	set := make(map[int]struct{})
	for i := 0; i < n; i++ {
		r := rand.Float64()
		for i, v := range cutoffs {
			if r <= v {
				set[i] = struct{}{}
				break
			}
		}
	}

	// Extract set's keys for selections
	sel := make([]int, len(set))
	i := 0
	for k, _ := range set {
		sel[i] = k
		i++
	}

	// Make a generation of just those chromosomes
	parents := &Generation{g.cherryPick(sel), nil}

	return parents
}
