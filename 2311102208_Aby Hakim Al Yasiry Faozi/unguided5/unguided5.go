package main

import "fmt"

func printOddNumbers(n int) {
	if n < 1 {
		return
	}
	printOddNumbers(n - 2)
	fmt.Print(n, " ")
}

func main() {
	var number int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&number)
	fmt.Printf("Barisan bilangan ganjil hingga %d: ", number)
	if number%2 == 0 {
		number--
	}
	printOddNumbers(number)
	fmt.Println()
}
