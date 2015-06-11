package ga

import "testing"

func TestExtractElite(t *testing.T) {
	tests := []struct {
		in       *Generation
		n        int
		perf     Performance
		outElite *Generation
		outRem   *Generation
	}{
		{
			&Generation{generation{
				&TestCM{0x1, 0x1, 0x1, 0x0},
				&TestCM{0x1, 0x0, 0x0, 0x0},
				&TestCM{0x1, 0x1, 0x0, 0x0},
			}, nil},
			1,
			&TestPerf{},
			&Generation{generation{
				&TestCM{0x1, 0x1, 0x1, 0x0},
			}, nil},
			&Generation{generation{
				&TestCM{0x1, 0x1, 0x0, 0x0},
				&TestCM{0x1, 0x0, 0x0, 0x0},
			}, nil},
		},
	}

	for num, test := range tests {
		outElite, outRem := test.in.extractElite(test.n, test.perf)
		if !generationsEqual(outElite.gen, test.outElite.gen) {
			t.Errorf("Test #%v: Elite not equal: expected %v, got %v",
				num+1, test.outElite.gen, outElite.gen)
		}
		if !generationsEqual(outRem.gen, test.outRem.gen) {
			t.Errorf("Test #%v: Rem not equal: expected %v, got %v",
				num+1, test.outRem.gen, outRem.gen)
		}
	}
}

func TestRoulette(t *testing.T) {
	t.Log("TODO")
}
