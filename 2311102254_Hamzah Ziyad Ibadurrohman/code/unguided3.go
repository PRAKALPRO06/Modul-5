package main

import "fmt"

func faktor_bilang(n, i int) {

	if i <= n {

		if n%i == 0 {
			fmt.Print(i, " ")
		}

		faktor_bilang(n, i+1)
	}
}
func main() {
	var n int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	faktor_bilang(n, 1)
}
