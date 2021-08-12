package main

import (
	"fmt"
)

func main() {
	fmt.Printf("the smallest positive number that is evenly divisible by all of the numbers from %v to %v is : %v\n", 1,20, smallestMultiple(1,20))
}

func smallestMultiple(from int, to int) int {
	for i := 1; i > 0; i++ {
		for j := from; j <= to; j++ {
			if i%j != 0 {
				break
			} else if j == 20 {
				return i
			}
		}
	}
	return 0
}