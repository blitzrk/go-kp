package ga

import "sort"

type generation []Chromosome

type Generation struct {
	generation
	pairs pairs
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

func (g *Generation) extractElite(p *Parameters) (*Generation, *Generation) {
	elite := make(generation, p.Pop)
	rem := make(generation, len(g.generation)-int(p.Pop))

	ps := make(pairs, len(g.generation))
	for i, c := range g.generation {
		ps[i] = pair{i, p.Fitness(c)}
	}
	sort.Sort(sort.Reverse(byScore(ps)))
	g.pairs = ps

	sortedChroms := g.cherryPick(ps.Items())
	copy(elite, sortedChroms[:p.Pop])
	copy(rem, sortedChroms[p.Pop:])

	return NewGeneration(elite), NewGeneration(rem)
}

func (g *Generation) roulette(p *Parameters) *Generation {
	// TODO(ben): implement roulette method
	return g
}

func (g1 *Generation) merge(g2 *Generation) *Generation {
	m := make(generation, len(g1.generation)+len(g2.generation))
	copy(m[:len(g1.generation)], g1.generation)
	copy(m[len(g1.generation):], g2.generation)
	return NewGeneration(m)
}

func (g *Generation) Select(p *Parameters) *Generation {
	elite, rem := g.extractElite(p)
	sel := rem.roulette(p)
	return sel.merge(elite)
}

func NewInitGen(p *Parameters) *Generation {
	gen := make(generation, p.Pop)

	var nGreedy int
	if gp, ok := interface{}(p).(GreedyAlg); ok {
		nGreedy = 1
		gen[len(gen)-1] = gp.Greedy()
	}

	for i := 0; i < len(gen)-nGreedy; i++ {
		gen[i] = NewRandChromosome(p)
	}

	return NewGeneration(gen)
}
