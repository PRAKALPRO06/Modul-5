package main

import "fmt"

// Fungsi rekursif untuk menghitung nilai suku ke-n dalam deret Fibonacci
func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int = 10 // Mengatur n ke 11 untuk mencetak hingga suku ke-10

	// Baris pertama: Suku ke-n
	fmt.Print("Suku n      : ")
	for i := 0; i <= n; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()

	// Baris kedua: Hasil Fibonacci
	fmt.Print("Fibonacci n : ")
	for i := 0; i <= n; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	fmt.Println()
}
