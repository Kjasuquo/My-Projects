package calculator

func Subtraction(x []float64) float64 {
	if len(x) == 0 {
		return 0
	}
	var ansSub = x[0]
	for i := 1; i < len(x); i++ {
		ansSub -= x[i]
	}
	return ansSub
}
