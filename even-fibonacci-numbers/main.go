package main

import (
	"fmt"
)

func main() {

	fibonacciSeq := []int{1, 2}
	sum := 2
	for  {
		size := len(fibonacciSeq)
		nextValue := fibonacciSeq[size-2] + fibonacciSeq[size-1]
		if nextValue > 4000000 {
			break
		} else if nextValue % 2 == 0 {
			sum += nextValue
		} 
		fibonacciSeq = append(fibonacciSeq, nextValue)
	}
	fmt.Printf("Some of even numbers = %v\n", sum)
}
