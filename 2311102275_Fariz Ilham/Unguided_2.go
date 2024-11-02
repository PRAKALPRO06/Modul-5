package main

import "fmt"

func pola(i, n int, s string) {
	if i <= n {
		s = s + "*"
		fmt.Println(s)
		pola(i+1, n, s)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan n: ")
	fmt.Scanln(&n)
	pola(1, n, "")
}
