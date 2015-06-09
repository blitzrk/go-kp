package ga

import (
	"math/rand"
	"time"
)

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

func (b *Breeder) Breed(ncross int, pcross, pmutate float64) *Generation {
	b.Gen = b.Gen.Cross(ncross, pcross)
	b.Mutate(pmutate)
	return b.Gen
}

func (g *Generation) Cross(n int, p float64) *Generation {
	return nil
}

func mutateChromosome(cm ChromosomeModel, p float64, r *rand.Rand) {
	for i := 0; i < cm.Len(); i++ {
		rv := r.Float64()
		if rv <= p {
			cm.MutateChar(i)
		}
	}
}

func (b *Breeder) Mutate(p float64) {
	for i := 0; i < len(b.Gen.gen); i++ {
		mutateChromosome(b.Gen.gen[i], p, b.r)
	}
}
