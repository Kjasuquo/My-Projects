package calculator

func Multiply(x []float64) float64 {
	if len(x) == 0 {
		return 0
	}
	var ansMult float64 = 1
	for i := 0; i < len(x); i++ {
		ansMult *= x[i]
	}
	return ansMult
}
