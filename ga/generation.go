package ga

import "sort"

type generation []Chromosome

type Generation struct {
	generation
	pairs pairs
}

func (g *Generation) rank(p Performance) pairs {
	if g.pairs != nil {
		return g.pairs
	}

	ps := make(pairs, len(g.generation))
	for i, c := range g.generation {
		ps[i] = pair{i, p.Fitness(c)}
	}
	sort.Sort(sort.Reverse(byScore(ps)))
	g.pairs = ps

	return ps
}

func NewGeneration(g generation) *Generation {
	return &Generation{g, nil}
}

func (g *Generation) cherryPick(its []int) generation {
	pick := make(generation, len(its))
	for i, v := range its {
		pick[i] = g.generation[v]
	}
	return pick
}

func (g *Generation) extractElite(n int, p Performance) (*Generation, *Generation) {
	elite := make(generation, n)
	rem := make(generation, len(g.generation)-n)

	sortedChroms := g.cherryPick(g.rank(p).Items())
	copy(elite, sortedChroms[:n])
	copy(rem, sortedChroms[n:])

	return NewGeneration(elite), NewGeneration(rem)
}

func (g *Generation) roulette(p Performance) *Generation {
	// TODO(ben): implement roulette method
	return g
}

func (g1 *Generation) merge(g2 *Generation) *Generation {
	m := make(generation, len(g1.generation)+len(g2.generation))
	copy(m[:len(g1.generation)], g1.generation)
	copy(m[len(g1.generation):], g2.generation)
	return &Generation{m, g1.pairs.MergeSortedDesc(g2.pairs)}
}

func (g *Generation) Select(nElite int, p Performance) *Generation {
	elite, rem := g.extractElite(nElite, p)
	sel := rem.roulette(p)
	return sel.merge(elite)
}

func NewInitGen(popSize int, randChrom func() Chromosome) *Generation {
	gen := make(generation, popSize)

	for i := 0; i < len(gen); i++ {
		gen[i] = randChrom()
	}

	return NewGeneration(gen)
}

type GreedyAlg interface {
	Greedy() Chromosome
}

func ImproveInitGen(gen *Generation, gp GreedyAlg) {
	gen.generation[0] = gp.Greedy()
}
