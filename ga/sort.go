package ga

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

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
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

	return merged
}
