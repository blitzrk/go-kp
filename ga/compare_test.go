package ga

func chromosomesEqual(c1, c2 ChromosomeModel) bool {
	if c1.Len() != c2.Len() {
		return false
	}
	for j := 0; j < c1.Len(); j++ {
		if c1.Loc(j) != c2.Loc(j) {
			return false
		}
	}
	return true
}

func generationsEqual(g1, g2 generation) bool {
	if len(g1) != len(g2) {
		return false
	}
	if len(g1) == 0 {
		return true
	}

	for i := 0; i < len(g1); i++ {
		if !chromosomesEqual(g1[i], g2[i]) {
			return false
		}
	}
	return true
}

func breedPairsEqual(bp1, bp2 []*breedPair) bool {
	if len(bp1) != len(bp2) {
		return false
	}
	for i := 0; i < len(bp1); i++ {
		if !chromosomesEqual(bp1[i].p1, bp2[i].p1) || !chromosomesEqual(bp1[i].p2, bp2[i].p2) {
			return false
		}
	}
	return true
}
