package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func printFibonacci(n, current int) {
	if current > n {
		return
	}
	fmt.Printf("%d ", fibonacci(current))
	printFibonacci(n, current+1)
}

func main() {
	printFibonacci(10, 0)
	fmt.Println()
}
