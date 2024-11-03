package main

import "fmt"

func faktor(n, i int) {
	if i <= n {
		if n%i == 0 {
			fmt.Print(i, " ")
			faktor(n, i+1)
		} else {
			faktor(n, i+1)
		}
	}
}

func main() {
	var n, i int

	i = 1

	fmt.Scanln(&n)

	faktor(n, i)
}
