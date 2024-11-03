package main

import (
	"fmt"
)

func fibonacci(term int) int {
	if term <= 1 {
		return term
	}
	return fibonacci(term-1) + fibonacci(term-2)
}

func main() {
	numTerms := 10
	fmt.Printf("Deret Fibonacci hingga suku ke-%d:\n", numTerms)
	for index := 0; index <= numTerms; index++ {
		fmt.Printf("Suku ke-%d: %d\n", index, fibonacci(index))
	}
}
