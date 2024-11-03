package main

import "fmt"

func fibonacci(n int) int {
	if n == 1 {
		return 1
	} else if n == 0 {
		return 0
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Masukkan Panjang n : ")
	fmt.Scan(&n)
	fmt.Print("n  : ")
	for i := 0; i < n; i++ {
		fmt.Print(i, " ")
	}

	fmt.Print("\nSn : ")
	for j := 0; j < n; j++ {
		fmt.Print(fibonacci(j), " ")
	}
}
