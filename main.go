package main

import (
	"fmt"
	"github.com/kjasuquo/calculate/calculator"
)

func main() {
	y := calculator.Calculate("1*2*3*4", "5+6+7+8", "256-21-41-72", "25/5/2")
	fmt.Println(y)
	fmt.Println(calculator.Multiply([]float64{}))
	fmt.Println(calculator.Addition([]float64{5, 6, 7, 8}))
	fmt.Println(calculator.Subtraction([]float64{256, 21, 41, 72}))
	fmt.Println(calculator.Division([]float64{25, 5, 0}))
}
