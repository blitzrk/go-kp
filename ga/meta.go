package ga

import "fmt"

type data struct {
	item  int
	score float64
}

type byScore []data

func (s byScore) Len() int           { return len(s) }
func (s byScore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byScore) Less(i, j int) bool { return s[i].score < s[j].score }

type metadata byScore

func (ps metadata) Items() (its []int) {
	its = make([]int, len(ps))
	for i, v := range ps {
		its[i] = v.item
	}
	return
}

func (ps metadata) Scores() (scs []float64) {
	scs = make([]float64, len(ps))
	for i, v := range ps {
		scs[i] = v.score
	}
	return
}

// Returns a new slice (replaces underlying array) of metadata with items
// renumbered. Thus it is required that the metadata already be in order.
func (ps metadata) Subset(i, j int) metadata {
	if j < i {
		panic(fmt.Sprintf("Invalid subset [%v, %v)", i, j))
	}

	sub := make(metadata, j-i)
	for k := 0; k < j-i; k++ {
		sub[k] = data{k, ps[i+k].score}
	}
	return sub
}

// Combines metadata by appending and then renumbers.
func (ps1 metadata) append(ps2 metadata) metadata {
	n1 := len(ps1)
	n2 := len(ps2)
	n := n1 + n2
	merged := make(metadata, n)
	copy(merged[:n1], ps1)
	copy(merged[n1:], ps2)

	for i := 0; i < n; i++ {
		merged[i].item = i
	}

	return merged
}
