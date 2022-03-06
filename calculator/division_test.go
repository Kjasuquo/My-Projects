package calculator

import (
	"math"
	"testing"
)

func TestDivision(t *testing.T) {
	var tests = []struct {
		input []float64
		want  float64
	}{
		{[]float64{250, 10, 2, 5}, 2.5},
		{[]float64{250, 10, 2, 0}, math.Inf(1)},
	}
	for _, test := range tests {
		ans := Division(test.input)
		if ans != test.want {
			t.Errorf("Expecting %v but got %v", test.want, ans)
		}
	}
}
