package calculator

import (
	"strconv"
	"strings"
)

func Calculate(s ...string) []float64 {
	var mainArray []float64
	var multi []float64
	var addit []float64
	var subtr []float64
	var divis []float64

	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "*") {
			mul := strings.Split(s[i], "*")
			for i := 0; i < len(mul); i++ {
				m, _ := strconv.Atoi(string(mul[i]))
				multi = append(multi, float64(m))
			}
			mainArray = append(mainArray, Multiply(multi))
		}

		if strings.Contains(s[i], "+") {

			add := strings.Split(s[i], "+")
			for i := 0; i < len(add); i++ {
				a, _ := strconv.Atoi(string(add[i]))
				addit = append(addit, float64(a))
			}
			mainArray = append(mainArray, Addition(addit))
		}

		if strings.Contains(s[i], "-") {
			sub := strings.Split(s[i], "-")
			for i := 0; i < len(sub); i++ {
				st, _ := strconv.Atoi(string(sub[i]))
				subtr = append(subtr, float64(st))
			}
			mainArray = append(mainArray, Subtraction(subtr))
		}

		if strings.Contains(s[i], "/") {
			div := strings.Split(s[i], "/")
			for i := 0; i < len(div); i++ {
				di, _ := strconv.Atoi(string(div[i]))
				divis = append(divis, float64(di))
			}
			mainArray = append(mainArray, Division(divis))
		}
	}
	return mainArray

}
