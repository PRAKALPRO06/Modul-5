package main

import "fmt"

func fibonacci(n int) int {
	var hasil int
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		hasil = fibonacci(n-1) + fibonacci(n-2)
		return hasil
	}
}

func main() {
	var n, hsl int
	fmt.Print("masukkan nilai suku nya: ")
	fmt.Scan(&n)
	hsl = fibonacci(n)
	fmt.Println("Hasilnya adalah: ", hsl)
}
