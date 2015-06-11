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
