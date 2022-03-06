package calculator

import "testing"

func TestMultiply(t *testing.T) {
	tests := []struct {
		input []float64
		want  float64
	}{
		{[]float64{0, 5, 7, 100}, 0},
		{[]float64{1, 1, 90}, 90},
		{[]float64{2, 80, 56, 7}, 62720},
	}
	//I decided to use the for loop iteration
	for i := 0; i < len(tests); i++ {
		ans := Multiply(tests[i].input)

		if ans != tests[i].want {
			t.Errorf("expected %v but getting %v", tests[i].want, ans)
		}
	}
}
