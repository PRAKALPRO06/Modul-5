package main

import "fmt"

func printSequence(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	printSequence(n - 1)
	fmt.Print(n, " ")
}

func main() {
	var number int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&number)
	fmt.Printf("Barisan bilangan: ")
	printSequence(number)
	fmt.Println()
}
