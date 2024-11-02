package main

import "fmt"

func ganjil(n int) {
	if n < 1 {
		return
	}

	if n%2 != 0 {
		ganjil(n - 2)
		fmt.Print(n, " ")
	} else {
		ganjil(n - 1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan n: ")
	fmt.Scanln(&n)
	ganjil(n)
}
