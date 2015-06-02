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
	merged := make(pairs, len(ps1)+len(ps2))
	n := minInt(len(ps1), len(ps2))

	for i := 0; i < n; i++ {
		if ps1[i].score > ps2[i].score {
			merged[i*2] = ps1[i]
			merged[i*2+1] = ps2[i]
		} else {
			merged[i*2] = ps2[i]
			merged[i*2+1] = ps1[i]
		}
	}

	if len(ps1) < len(ps2) {
		copy(merged[n*2:], ps2[n:])
	} else {
		copy(merged[n*2:], ps1[n:])
	}
	return merged
}
