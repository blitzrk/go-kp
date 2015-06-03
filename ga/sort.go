package ga

import "fmt"

type pair struct {
	item  int
	score float64
}

type byScore []pair

func (s byScore) Len() int           { return len(s) }
func (s byScore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byScore) Less(i, j int) bool { return s[i].score < s[j].score }

type pairs byScore

func (ps pairs) Items() (its []int) {
	its = make([]int, len(ps))
	for i, v := range ps {
		its[i] = v.item
	}
	return
}

func (ps pairs) Scores() (scs []float64) {
	scs = make([]float64, len(ps))
	for i, v := range ps {
		scs[i] = v.score
	}
	return
}

func (ps pairs) Renumber() {
	for i, p := range ps {
		p.item = i
	}
}

func (ps pairs) Subset(i, j int) pairs {
	if j < i {
		panic(fmt.Sprintf("Invalid subset [%v, %v)", i, j))
	}

	sub := make(pairs, j-i)
	for k := 0; k < j-i; k++ {
		sub[k] = pair{k + 1, ps[i+k].score}
	}
	sub.Renumber()
	return sub
}

func (ps1 pairs) MergeSortedDesc(ps2 pairs) pairs {
	n1 := len(ps1)
	n2 := len(ps2)
	n := n1 + n2
	merged := make(pairs, n)

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
	}

	merged.Renumber()
	return merged
}
