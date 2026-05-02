package main

import (
	"fmt"
)

func generate() <-chan int {
	out := make(chan int, 10)
	
	go func() {
		defer close(out)
		
		for i := 1; i <= 100; i++ {
			out <- i
		}
	}()
	
	return out
}

func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	
	go func() {
		defer close(out)
		
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
	}()
	
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int, 5)
	
	go func() {
		defer close(out)
		
		for n := range in {
			out <- n * n
		}
	}()
	
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)
	
	go func() {
		defer close(out)
		
		total := 0
		for n := range in {
			total += n
		}
		out <- total
	}()
	
	return out
}

func main() {
	

	stage1 := generate()         
	stage2 := filterEven(stage1) 
	stage3 := square(stage2)     
	stage4 := sum(stage3)        

	finalResult := <- stage4

	fmt.Printf("result: %d\n", finalResult)
}