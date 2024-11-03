package main

import "fmt"

func tampilGanjil(n int, i int) {
	if i > n {
		return
	}
	fmt.Print(i, " ")
	tampilGanjil(n, i+2)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	fmt.Print("Bilangan ganjil dari 1 hingga ", n, ": ")
	tampilGanjil(n, 1)
	fmt.Println()
}
