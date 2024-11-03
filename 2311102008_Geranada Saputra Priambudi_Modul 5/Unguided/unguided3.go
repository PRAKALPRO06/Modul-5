package main

import "fmt"

func faktorBilangan(n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktorBilangan(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)

	fmt.Print("Faktor dari ", n, ": ")
	faktorBilangan(n, 1)
	fmt.Println()
}
