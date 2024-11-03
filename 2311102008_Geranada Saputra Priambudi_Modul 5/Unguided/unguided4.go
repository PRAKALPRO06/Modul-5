package main

import "fmt"

func barisan(n int, current int) {
	if current < 1 {
		return
	}
	fmt.Print(current, " ")
	barisan(n, current-1)
	fmt.Print(current, " ")
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)

	fmt.Print("Hasil: ")
	barisan(n, n)
	fmt.Println()
}
