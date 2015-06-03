package ga

import (
	"bytes"
	"fmt"
	"math/rand"
)

type Chromosome []byte

type Performance interface {
	Fitness(Chromosome) float64
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

type Parameters struct {
	Performance
	ChromLen   uint
	ChromChars uint8
	Pop        uint
	Elite      uint8
	Crosses    uint8
	CrossProb  float64
	MutateProb float64
	MaxGens    uint32
}

func randChromosomeFunc(chars, clen int) func() Chromosome {
	return func() Chromosome {
		c := make(Chromosome, clen)
		for i := 0; i < len(c); i++ {
			c[i] = byte(rand.Intn(chars))
		}
		return c
	}
}

func Run(p *Parameters) error {
	for gen := 0; gen < int(p.MaxGens); gen++ {
		// Generate initial population
		cf := randChromosomeFunc(int(p.ChromChars), int(p.ChromLen))
		g := NewInitGen(int(p.Pop), cf)
		if gp, ok := interface{}(p).(GreedyPerformance); ok {
			ImproveInitGen(g, gp)
		}

		// Select portion of population to breed
		breeders := g.Select(int(p.Elite), p.Performance)

		// TODO(ben): Breed
		_ = breeders
	}
	return nil
}
