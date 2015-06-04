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

// Takes two individually ordered data slices and sorts them into a new slice
// in O(n+m) time. It assumes that the related generation is merged with the
// second appended to the first and renumbers as such.
func (ps1 metadata) MergeSortedDesc(ps2 metadata) metadata {
	n1 := len(ps1)
	n2 := len(ps2)
	n := n1 + n2
	merged := make(metadata, n)

	var j1, j2 int
	for i := 0; i < n; i++ {
		if j1 >= n1 {
			copy(merged[i:], ps2[j2:])
			break
		} else if j2 >= n2 {
			copy(merged[i:], ps1[j1:])
			break
		}

		if ps1[j1].score > ps2[j2].score {
			merged[i] = ps1[j1]
			j1++
		} else {
			merged[i] = ps2[j2]
			j2++
		}

		// Renumbers metadata
		merged[i].item = i
	}

	return merged
}
