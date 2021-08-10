package main

import (
	"fmt"
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
		fmt.Printf("The sum of all the multiples of 3 or 5 below %v is %v.\n",n ,multiplesOf3Or5(n))
	}
}

func multiplesOf3Or5(n int) int {
	var sum = 0
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
