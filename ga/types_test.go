package ga

type TestCM Chromosome

func (cm *TestCM) Key() string      { return string(*cm) }
func (cm *TestCM) Len() int         { return len(*cm) }
func (cm *TestCM) Loc(i int) byte   { return (*cm)[i] }
func (cm *TestCM) MutateChar(i int) { return }
func (cm *TestCM) String() string   { return Chromosome(*cm).String() }

func (cm1 *TestCM) Cross(locus int, cm2 ChromosomeModel) (ChromosomeModel, ChromosomeModel) {
	return cm1, cm2
}

type TestCMMut struct {
	*TestCM
}

func (cm TestCMMut) MutateChar(i int) {
	(*cm.TestCM)[i] = 1 - (*cm.TestCM)[i]
}

func (cm1 TestCMMut) Cross(locus int, cm2 ChromosomeModel) (ChromosomeModel, ChromosomeModel) {
	return nil, nil
}

type TestPerf struct {
	Length int
}

func (p *TestPerf) Fitness(cm ChromosomeModel) float64 {
	var sum float64
	for i := 0; i < cm.Len(); i++ {
		sum += float64(int(cm.Loc(i)))
	}
	return sum
}

func (p *TestPerf) Rand() ChromosomeModel {
	cm := make(TestCM, p.Length)
	return &cm
}

func (p *TestPerf) Greedy() ChromosomeModel {
	cm := make(TestCM, p.Length)
	if p.Length > 0 {
		cm[0] = 0x1
	}
	return &cm
}

// Source that alternates between 0 and 0.999...
type TestRandSource struct {
	last int64
}

// const maxRandInt63 int64 = 1<<62 | 1<<61 | 1<<60 | 1<<59 | 1<<58 | 1<<57 | 1<<56 | 1<<55 | 1<<54 | 1<<53 | 1<<52 | 1<<51
const maxRandInt63 int64 = 1<<63 - 1<<10

func (s *TestRandSource) Seed(seed int64) {}
func (s *TestRandSource) Int63() int64 {
	if s.last == 0 {
		s.last = maxRandInt63
		return maxRandInt63
	} else {
		s.last = 0
		return 0
	}
}
