package main

import "fmt"

func tampilBilangan(n int, m int) {
	if m < 1 {
		return
	}
	fmt.Print(m, " ")
	tampilBilangan(n, m-1)
	fmt.Print(m, " ")
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	fmt.Print("Hasil: ")
	tampilBilangan(n, n)
	fmt.Println()
}
