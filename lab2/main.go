package main

import (
	"fmt"
)

type Currency interface {
	Name() string
	USDRate() float64
	IsCrypto() bool
}



type USD struct {
	rate float64 
}

func (u USD) Name() string {
	return "US Dollar"
}

func (u USD) USDRate() float64 {
	return u.rate
}

func (u USD) IsCrypto() bool {
	return false
}



type Bitcoin struct {
	rate float64 
}

func (b Bitcoin) Name() string {
	return "Bitcoin"
}

func (b Bitcoin) USDRate() float64 {
	return b.rate
}

func (b Bitcoin) IsCrypto() bool {
	return true
}



func main() {

	b := []int{10, 25, 42, 8, 99, 14, 5, 33, 7, 50}
	a := [10]int{1,2,3,4,5,6,7,8,9,10}
	
	result := make([]int, 10)

	for i := 0; i < 10; i++ {
		result[i] = a[i] + b[i]
	}

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("result:", result)
	fmt.Println()



	myUSD := USD{rate: 1.0}
	myBTC := Bitcoin{rate: 64500.50}

	currencies := []Currency{myUSD, myBTC}

	for _, c := range currencies {
		fmt.Printf("money: %s\n", c.Name())
		fmt.Printf("money to USD: %.2f\n", c.USDRate())
		fmt.Printf("is crypto? %t\n", c.IsCrypto())
	}
}