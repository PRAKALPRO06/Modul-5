package main

import "fmt"

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * faktorial(n-1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan non-negatif: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Faktorial tidak didefinisikan untuk bilangan negatif")
	} else {
		fmt.Println("Hasil faktorial:", faktorial(n))
	}
}
