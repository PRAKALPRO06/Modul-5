package main

import "fmt"

func power(base, exponent int) int {
	if exponent == 0 {
		return 1
	} else if exponent < 0 {
		return 1 / power(base, -exponent)
	} else {
		return base * power(base, exponent-1)
	}
}

func main() {
	var base, exponent int
	fmt.Print("Masukkan bilangan pokok: ")
	fmt.Scan(&base)
	fmt.Print("Masukkan pangkat: ")
	fmt.Scan(&exponent)

	result := power(base, exponent)
	fmt.Printf("%d pangkat %d adalah %d\n", base, exponent, result)
}
