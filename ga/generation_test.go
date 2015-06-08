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

func TestCherryPick(t *testing.T) {

}
