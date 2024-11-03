package main

import "fmt"

func printFactors(n, divisor int) {
	if divisor > n {
		return
	}
	if n%divisor == 0 {
		fmt.Printf("%d ", divisor)
	}
	printFactors(n, divisor+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&n)

	fmt.Printf("Faktor dari %d adalah: ", n)
	printFactors(n, 1)
	fmt.Println()
}
