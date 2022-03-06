package calculator

func Division(x []float64) float64 {
	if len(x) == 0 {
		return 0
	}
	var ansDiv = x[0]
	for i := 1; i < len(x); i++ {
		ansDiv /= x[i]
	}
	return ansDiv
}
