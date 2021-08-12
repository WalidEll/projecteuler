package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Println("Please enter a value!")
	} else {
		n, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Printf("The difference between the sum of the squares of the first %v natural numbers and the square of the sum is %v.\n",n ,int(sumSquareDifference(n)))
	}
}

func sumSquareDifference(n int)  float64 {
	sumOfSquares := 0.
	sums := 0.
	for i := 1; i <= n; i++ {
		sumOfSquares += math.Pow(float64(i), 2)
	}
	for i := 1; i <= n; i++ {
		sums += float64(i);
	}
	squareOfSums := math.Pow(sums, 2)
	fmt.Println(squareOfSums, sumOfSquares)
	return  squareOfSums - sumOfSquares
}