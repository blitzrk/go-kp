package ga

import "testing"

type TestCM []byte

func (cm *TestCM) Key() string      { return string(*cm) }
func (cm *TestCM) Len() int         { return len(*cm) }
func (cm *TestCM) Loc(i int) byte   { return (*cm)[i] }
func (cm *TestCM) MutateChar(i int) { return }

type TestPerf struct{}

func (p *TestPerf) Fitness(cm ChromosomeModel) float64 {
	var sum float64
	for i := 0; i < cm.Len(); i++ {
		sum += float64(int(cm.Loc(i)))
	}
	return sum
}

func (p *TestPerf) Rand(i int) ChromosomeModel {
	cm := make(TestCM, i)
	return &cm
}

func TestRank(t *testing.T) {
	tests := []struct {
		in   *Generation
		perf Performance
		out  metadata
	}{
		{
			&Generation{
				generation{
					&TestCM{0x1, 0x1, 0x1, 0x0},
					&TestCM{0x1, 0x1, 0x0, 0x0},
					&TestCM{0x1, 0x1, 0x1, 0x1},
					&TestCM{0x1, 0x0, 0x0, 0x0},
				},
				nil,
			},
			&TestPerf{},
			metadata{
				data{2, 4},
				data{0, 3},
				data{1, 2},
				data{3, 1},
			},
		},
	}

	for num, test := range tests {
		out := test.in.rank(test.perf)
		if !metadataEqual(out, test.out) {
			t.Errorf("Test #%v failed: Expected %v, got %v.\n", num+1, test.out, out)
		}
	}
}

func generationsEqual(g1, g2 generation) bool {
	if len(g1) != len(g2) {
		return false
	}
	if len(g1) == 0 {
		return true
	}

	for i := 0; i < len(g1); i++ {
		if g1[i].Len() != g2[i].Len() {
			return false
		}
		for j := 0; j < g1[i].Len(); j++ {
			if g1[i].Loc(j) != g2[i].Loc(j) {
				return false
			}
		}
	}
	return true
}

func TestGenerationCherryPick(t *testing.T) {
	tests := []struct {
		gen *Generation
		in  []int
		out generation
	}{
		{
			&Generation{
				generation{
					&TestCM{0x1, 0x1, 0x1, 0x0},
					&TestCM{0x1, 0x1, 0x0, 0x0},
					&TestCM{0x1, 0x1, 0x1, 0x1},
					&TestCM{0x1, 0x0, 0x0, 0x0},
				},
				nil,
			},
			[]int{1, 0, 2},
			generation{
				&TestCM{0x1, 0x1, 0x0, 0x0},
				&TestCM{0x1, 0x1, 0x1, 0x0},
				&TestCM{0x1, 0x1, 0x1, 0x1},
			},
		},
	}

	for num, test := range tests {
		out := test.gen.cherryPick(test.in)
		if !generationsEqual(out, test.out) {
			t.Errorf("Test #%v failed: Expected %v, got %v.\n", num+1, test.out, out)
		}
	}
}

func TestGenerationMerge(t *testing.T) {
	tests := []struct {
		in1 *Generation
		in2 *Generation
		out *Generation
	}{
		{
			&Generation{
				generation{
					&TestCM{0x1, 0x1, 0x1, 0x0},
					&TestCM{0x1, 0x1, 0x0, 0x0},
				},
				nil,
			},
			&Generation{
				generation{
					&TestCM{0x1, 0x1, 0x1, 0x1},
					&TestCM{0x1, 0x0, 0x0, 0x0},
				},
				nil,
			},
			&Generation{
				generation{
					&TestCM{0x1, 0x1, 0x1, 0x0},
					&TestCM{0x1, 0x1, 0x0, 0x0},
					&TestCM{0x1, 0x1, 0x1, 0x1},
					&TestCM{0x1, 0x0, 0x0, 0x0},
				},
				nil,
			},
		},
	}

	for num, test := range tests {
		out := test.in1.merge(test.in2)
		if !generationsEqual(out.gen, test.out.gen) || !metadataEqual(out.meta, test.out.meta) {
			t.Errorf("Test #%v failed: Expected %v, got %v.\n", num+1, test.out, out)
		}
	}
}
