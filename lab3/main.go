package main

import (
	"Lab3/calc" 
	"fmt"
)

func main() {
	

	fmt.Println("Сума (1.5, 2.5, 6.0) =", calc.Sum(1.5, 2.5, 6.0))
	fmt.Println("Максимум (4.2, 9.1, -2.3) =", calc.Max(4.2, 9.1, -2.3))
	fmt.Println("Мінімум (4.2, 9.1, -2.3) =", calc.Min(4.2, 9.1, -2.3))

	res, err := calc.Divide(10.0, 2.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ділення (10 / 2) =", res)
	}

	_, err = calc.Divide(10.0, 0.0) 
	if err != nil {
		fmt.Println("Спроба ділення на 0:", err)
	}


	var myCalculator calc.Calculator = calc.Calc{}

	fmt.Println("Calc Сума (10, 20, 30) =", myCalculator.Sum(10, 20, 30))
	fmt.Println("Calc Максимум (100, 50) =", myCalculator.Max(100, 50))
	fmt.Println("Calc Мінімум (100, 10) =", myCalculator.Min(100, 10))

	divRes, divErr := myCalculator.Divide(25, 5)
	if divErr == nil {
		fmt.Println("Calc Ділення (25 / 5) =", divRes)
	}
}
