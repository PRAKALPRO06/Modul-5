package main

import "fmt"

func bilanganGanjil(n int, current int) {
	if current > n {
		return
	}
	fmt.Print(current, " ")
	bilanganGanjil(n, current+2)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)

	fmt.Print("Bilangan ganjil dari 1 hingga ", n, ": ")
	bilanganGanjil(n, 1)
	fmt.Println()
}
