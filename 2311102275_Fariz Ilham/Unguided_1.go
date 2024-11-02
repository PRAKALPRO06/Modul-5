package main

import "fmt"

func fibonacci(suku0, suku1, n int) int {
	if n == 0 {
		return suku0
	} else if n == 1 {
		return suku1
	} else {
		return fibonacci(suku1, suku0+suku1, n-1)
	}
}

func main() {
	var suku0, suku1 int

	fmt.Print("Masukkan nilai suku ke-0: ")
	fmt.Scanln(&suku0)
	fmt.Print("Masukkan nilai suku ke-1: ")
	fmt.Scanln(&suku1)

	fmt.Print("n = ")
	for n := 0; n <= 10; n++ {
		fmt.Print(n, " ")
	}
	fmt.Println()

	fmt.Print("Sn = ")
	for n := 0; n <= 10; n++ {
		fmt.Print(fibonacci(suku0, suku1, n), " ")
	}
	fmt.Println()
}
