package ga

import (
	"fmt"
	"math/rand"
)

type Chromosome []byte

type Performance interface {
	Fitness(Chromosome) float64
}

func (c Chromosome) String() string {
	return fmt.Sprintf("%v", []uint8(c))
}

type generation []Chromosome

type Parameters struct {
	Performance
	ChromLen   uint
	Pop        uint
	Elite      uint8
	Crosses    uint8
	CrossProb  float64
	MutateProb float64
	MaxGens    uint32
}

func NewRandChromosome(p *Parameters) Chromosome {
	c := make(Chromosome, p.ChromLen)
	for i := 0; i < len(c); i++ {
		c[i] = byte(rand.Intn(2))
	}
	return c
}

type Greedyer interface {
	Greedy() Chromosome
}

func NewRandPop(p *Parameters) generation {
	gen := make(generation, p.Pop)

	var nGreedy int
	if gp, ok := interface{}(p).(Greedyer); ok {
		nGreedy = 1
		gen[len(gen)-1] = gp.Greedy()
	}

	for i := 0; i < len(gen)-nGreedy; i++ {
		gen[i] = NewRandChromosome(p)
	}

	return gen
}
