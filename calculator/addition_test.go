package calculator

import "testing"

func TestAddition(t *testing.T) {
	var tests = []struct {
		input []float64
		want  float64
	}{
		{[]float64{1, -1, 3, 4}, float64(7)},
		{[]float64{6, 8, 8, 0, 9}, float64(31)},
	}
	//using a range loop to iterate through
	for _, test := range tests {
		got := Addition(test.input)
		if got != test.want {
			t.Errorf("Expected %v but got %v", test.want, got)
		}
	}
}
