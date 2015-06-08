package ga

import "math/rand"

func (g *Generation) Breed(ncross int, pcross, pmutate float64) *Generation {
	children := g.Cross(ncross, pcross)
	children.Mutate(pmutate)
	return children
}

func (g *Generation) Cross(n int, p float64) *Generation {
	return nil
}

func MutateChromosome(cm ChromosomeModel, p float64) {
	for i := 0; i < cm.Len(); i++ {
		r := rand.Float64()
		if r <= p {
			cm.MutateChar(i)
		}
	}
}

func (g *Generation) Mutate(p float64) {
	for i := 0; i < len(g.gen); i++ {
		MutateChromosome(g.gen[i], p)
	}
}
