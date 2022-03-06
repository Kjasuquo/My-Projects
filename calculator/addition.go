package calculator

func Addition(x []float64) float64 {
	var ansAdd float64 = 0
	for i := 0; i < len(x); i++ {
		ansAdd += x[i]
	}
	return ansAdd
}
