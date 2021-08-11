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
		fmt.Printf("The largest prime factor of %v is %v\n",n ,largestPrimeFactor(n))
	}
}

func largestPrimeFactor(n int) int {
	largestPrimeFactor := -1
	d := 2
	for n != 0 {
		if n % d == 0 {
			largestPrimeFactor = d
			n /= d
			if n == 1 {
				break
			}
		} else {
			d++
		}
	}
	return largestPrimeFactor
}
