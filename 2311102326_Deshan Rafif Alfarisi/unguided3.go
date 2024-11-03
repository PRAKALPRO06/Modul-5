package main

import "fmt"

func findFactors(num, divisor int) {
	if divisor > num {
		return
	}
	if num%divisor == 0 {
		fmt.Print(divisor, " ")
	}
	findFactors(num, divisor+1)
}

func main() {
	var number int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&number)
	fmt.Printf("Faktor dari %d adalah: ", number)
	findFactors(number, 1)
	fmt.Println()
}
