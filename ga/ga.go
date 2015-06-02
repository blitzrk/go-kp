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

func NewRandChromosome(p *Parameters) Chromosome {
	c := make(Chromosome, p.ChromLen)
	for i := 0; i < len(c); i++ {
		c[i] = byte(rand.Intn(int(p.ChromChars)))
	}
	return c
}

type GreedyAlg interface {
	Greedy() Chromosome
}
