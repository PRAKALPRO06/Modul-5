package main

import "fmt"

func power(base, exponent int) int {
	if exponent == 0 {
		return 1
	} else if exponent == 1 {
		return base
	} else {
		return base * power(base, exponent-1)
	}
}

func main() {
	var base, exponent int
	fmt.Print("Masukkan nilai basis (base): ")
	fmt.Scanln(&base)
	fmt.Print("Masukkan nilai eksponen (exponent): ")
	fmt.Scanln(&exponent)

	fmt.Printf("%d pangkat %d adalah %d\n", base, exponent, power(base, exponent))
}
