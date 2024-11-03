package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("Deret Fibonacci hingga suku ke-10:")
	fmt.Println("n  : 0  1  2  3  4  5  6   7   8   9   10")
	fmt.Print("Sn : ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%-3d", fibonacci(i))
	}
	fmt.Println()
}
