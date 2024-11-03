package main

import "fmt"

func fibonacci(n, skrng, cetak, next int) int {
	if skrng >= n {
		fmt.Println()
		return 0
	}
	fmt.Print(cetak, " ")
	return fibonacci(n, skrng+1, next, cetak+next)
}

func main() {
	var n int
	fmt.Scan(&n)
	fibonacci(n, 0, 0, 1)
}
