package main

import "fmt"

// fungsi rekursif untuk menghitung nilai Fibonacci ke-n
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// contoh penggunaan fungsi fibonacci
	for i := 0; i <= 10; i++ {
		fmt.Printf("Fibonacci ke-%d: %d\n", i, fibonacci(i))
	}
}
