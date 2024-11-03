// Meutya Azzahra Efendi
// 2311102166
// IF-11-06
package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n (indeks deret Fibonacci): ")
	fmt.Scanln(&n)

	hasil := fibonacci(n)
	fmt.Printf("Nilai Fibonacci ke-%d adalah: %d\n", n, hasil)
}
