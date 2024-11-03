package main

import "fmt"

func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("  %d  ", i)
	}
	fmt.Println()
	for i := 0; i <= 10; i++ {
		fmt.Printf("  %d  ", fibo(i))
	}
}
