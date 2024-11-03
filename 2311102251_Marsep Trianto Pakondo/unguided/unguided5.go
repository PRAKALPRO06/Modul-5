package main

import "fmt"

func bilanganGanjil(n, x int)  {
	if x == n {
		if x % 2 != 0 {
			fmt.Print(x)
		}
	} else {
		if x % 2 != 0 {
			fmt.Print(x, " ")
			bilanganGanjil(n, x+1)
		} else {
			bilanganGanjil(n, x+1)
		}
	}
	
}

func main()  {
	var n int

	fmt.Print("Masukkan bilangan N : ")
	fmt.Scan(&n)

	bilanganGanjil(n, 1)
}