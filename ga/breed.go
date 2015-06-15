package ga

import (
	"fmt"
	"math/rand"
	"time"
)

type breedPair struct {
	p1, p2 ChromosomeModel
}

func (bp *breedPair) String() string {
	return fmt.Sprintf("[%v, %v]", bp.p1, bp.p2)
}

func (bp *breedPair) Crossover(locus int) {
	bp.p1, bp.p2 = bp.p1.Cross(locus, bp.p2)
}

func (bp *breedPair) Children() []ChromosomeModel {
	return []ChromosomeModel{bp.p1, bp.p2}
}

func (bp *breedPair) Len() int {
	if bp.p1.Len() != bp.p2.Len() {
		panic("Chromosomes in same generation of unequal size!")
	}
	return bp.p1.Len()
}

type Breeder struct {
	r   *rand.Rand
	Gen *Generation
}

func NewBreeder(g *Generation) *Breeder {
	return &Breeder{
		r:   rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
		Gen: g,
	}
}

func (b *Breeder) pairOff() []*breedPair {
	g := b.Gen.gen
	numPairs := len(g) / 2
	pairs := make([]*breedPair, numPairs)

	for i := 0; i < numPairs; i++ {
		pairs[i] = &breedPair{g[i*2], g[i*2+1]}
	}

	// One parent is polygamous if uneven number...
	if len(g)%2 == 1 {
		pairs = append(pairs, &breedPair{g[len(g)-2], g[len(g)-1]})
	}

	return pairs
}

func (b *Breeder) Breed(ncross int, pcross, pmutate float64) *Generation {
	parents := b.pairOff()
	children := make(generation, 0)

	for _, pair := range parents {
		if rv := b.r.Float64(); rv < pcross {
			for range [2]struct{}{} {
				locus := b.r.Intn(pair.Len())
				pair.Crossover(locus)
			}
		}
		children = append(children, pair.Children()...)
	}
	b.Gen = &Generation{children, nil}

	b.mutate(pmutate)
	return b.Gen
}

func mutateChromosome(cm ChromosomeModel, p float64, r *rand.Rand) {
	for i := 0; i < cm.Len(); i++ {
		if rv := r.Float64(); rv <= p {
			cm.MutateChar(i)
		}
	}
}

func (b *Breeder) mutate(p float64) {
	for i := 0; i < len(b.Gen.gen); i++ {
		mutateChromosome(b.Gen.gen[i], p, b.r)
	}
}
