package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int = 11
	fmt.Println("Deret Fibonacci hingga suku ke-11:")

	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}