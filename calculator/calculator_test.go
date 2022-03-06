package calculator

import (
	"reflect"
	"testing"
)

/*
kit
keyword, name type
*/
func TestCalculate(t *testing.T) {
	var tests = []struct {
		input []string
		want  []float64
	}{
		{[]string{"1*2*3*4", "5+6+7+8", "0-5-3", "25/5/2"}, []float64{24, 26, -8, 2.5}},
		{[]string{"5*4*3*2", "2/3/4/5", "1+2+3+4", "8-2-4-1"}, []float64{120, 0.03333333333333333, 10, 1}},
		{[]string{"1*2*3*4", "20/4/2/1", "4+3+2+1", "4-3-2-1"}, []float64{24, 2.5, 10, -2}},
		{[]string{"5*2*3*5", "30/2/5/1", "6+6+3+3", "41-21-10-5"}, []float64{150, 3, 18, 5}},
		{[]string{"4*4*4*4", "100/5/10/2", "6+6+3+3", "64-24-10-11"}, []float64{256, 1, 18, 19}},
		{[]string{"6+6+3+3", "64-24-10-11"}, []float64{18, 19}},
		{[]string{"6*6*3*3", "64/2/8/2"}, []float64{324, 2}},
		{[]string{"10+11+12+13", "11-2-4-4", "10*4*2*1", "1000/10/5/5"}, []float64{46, 1, 80, 4}},
		{[]string{"5000/10/10/10"}, []float64{5}},
		{[]string{"1001+2+3+4"}, []float64{1010}},
		{[]string{"1*2*3*4", "4+3+2+1", "4-3-2-1"}, []float64{24, 10, -2}},
	}
	for _, test := range tests {
		ans := Calculate(test.input...)
		if reflect.DeepEqual(ans, test.want) == false {
			t.Errorf("Expected %v but getting %v", test.want, ans)
		}
	}
}
