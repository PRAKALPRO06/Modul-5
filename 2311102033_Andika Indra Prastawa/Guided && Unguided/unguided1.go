package main

import "fmt"

// Fungsi rekursif untuk menghitung nilai Fibonacci ke-n
func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n untuk menghitung deret Fibonacci hingga suku ke-n: ")
	fmt.Scanln(&n)
	fmt.Println("Deret Fibonacci hingga suku ke-", n, ":")

	// Menampilkan deret Fibonacci hingga suku ke-n
	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
	fmt.Println()
}
