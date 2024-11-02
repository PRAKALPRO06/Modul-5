package main

import "fmt"

func printDescending(n int) {
	if n < 1 {
		return
	}
	fmt.Print(n, " ")
	printDescending(n - 1)
}

func printAscending(start, n int) {
	if start > n {
		return
	}
	fmt.Print(start, " ")
	printAscending(start+1, n)
}

func printPattern(n int) {
	printDescending(n)
	printAscending(2, n)
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scanln(&n)

	printPattern(n)
}
