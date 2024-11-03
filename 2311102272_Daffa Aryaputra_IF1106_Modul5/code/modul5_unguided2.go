package main

import "fmt"

func printStars(n int) {
	if n > 0 {
		fmt.Print("*")
		printStars(n - 1)
	}
}

func printPattern(n int) {
	if n > 0 {
		printPattern(n - 1)
		printStars(n)
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n)

	fmt.Println("Pola bintang:")
	printPattern(n)
}
