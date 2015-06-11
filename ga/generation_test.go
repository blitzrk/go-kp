package ga

import "testing"

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
		{
			&Generation{
				generation{
					&TestCM{0x1, 0x1, 0x1, 0x1},
					&TestCM{0x1, 0x1, 0x1, 0x0},
				},
				metadata{
					data{0, 4},
					data{1, 3},
				},
			},
			&Generation{
				generation{
					&TestCM{0x1, 0x1, 0x0, 0x0},
					&TestCM{0x1, 0x0, 0x0, 0x0},
				},
				metadata{
					data{0, 2},
					data{1, 1},
				},
			},
			&Generation{
				generation{
					&TestCM{0x1, 0x1, 0x1, 0x1},
					&TestCM{0x1, 0x1, 0x1, 0x0},
					&TestCM{0x1, 0x1, 0x0, 0x0},
					&TestCM{0x1, 0x0, 0x0, 0x0},
				},
				metadata{
					data{0, 4},
					data{1, 3},
					data{2, 2},
					data{3, 1},
				},
			},
		},
	}

	for num, test := range tests {
		out := test.in1.append(test.in2)
		if !generationsEqual(out.gen, test.out.gen) || !metadataEqual(out.meta, test.out.meta) {
			t.Errorf("Test #%v failed: Expected %v, got %v.\n", num+1, test.out, out)
		}
	}
}

func TestImproveInitGen(t *testing.T) {
	tests := []struct {
		in   *Generation
		perf GreedyPerformance
		out  *Generation
	}{
		{
			&Generation{
				generation{
					&TestCM{0x0, 0x0, 0x0, 0x0},
					&TestCM{0x0, 0x0, 0x0, 0x0},
				},
				nil,
			},
			&TestPerf{4},
			&Generation{
				generation{
					&TestCM{0x0, 0x0, 0x0, 0x0},
					&TestCM{0x1, 0x0, 0x0, 0x0},
				},
				metadata{
					data{1, 1},
					data{0, 0},
				},
			},
		},
	}

	for num, test := range tests {
		out := &Generation{make(generation, len(test.in.gen)), nil}
		copy(out.gen, test.in.gen)
		if test.in.meta != nil {
			out.meta = make(metadata, len(test.in.meta))
			copy(out.meta, test.in.meta)
		}

		ImproveInitGen(out, test.perf)
		if !generationsEqual(out.gen, test.out.gen) || !metadataEqual(out.meta, test.out.meta) {
			t.Errorf("Test #%v failed: Expected %v, got %v.\n", num+1, test.out, out)
		}
	}
}
