package main

import (
	"fmt"
)

func Fibonacci(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return Fibonacci(x-1) + Fibonacci(x-2)
	}
}

func CetakFibonacci(x int) {
	if x >= 0 {
		CetakFibonacci(x - 1)
		fmt.Println(Fibonacci(x))
	}
}

func main() {
	var x int
	fmt.Print("Masukkan angka untuk menampilkan deret Fibonacci hingga suku ke-n: ")
	fmt.Scanln(&x)

	fmt.Printf("Deret Fibonacci hingga suku ke-%d:\n", x)
	CetakFibonacci(x)
}
