package main

import "fmt"

func pola(i, N int, b string) {
	if i <= N {
		b = "*" + b
		fmt.Println(b)
		pola(i+1, N, b)
	}
}

func main() {
	var N int
	var b string

	fmt.Scanln(&N)
	pola(1, N, b)
}
