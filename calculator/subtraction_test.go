package calculator

import "testing"

func TestSubtraction(t *testing.T) {
	tests := []struct {
		input []float64
		want  float64
	}{
		{[]float64{26, 8, 8, 1, 9}, float64(0)},
		{[]float64{0, 10, 7, 6.9, 8.1}, -32},
		{[]float64{1, 2, 3, 4}, -8},
		{[]float64{0, 0, 0}, 0},
	}
	//using a range loop to iterate through
	for _, test := range tests {
		ans := Subtraction(test.input)
		if ans != test.want {
			t.Errorf("expected %v but got %v", test.want, ans)
		}
	}
}
