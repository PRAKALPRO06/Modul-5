package main

import "fmt"

func faktor(n, x int)  {
	if x == n {
		fmt.Print(x)
	} else {
		if n % x == 0 {
			fmt.Print(x, " ")
			faktor(n, x+1)
		} else {		
			faktor(n, x+1)
		}
	}
}

func main()  {
	var n int

	fmt.Print("Masukkan bilangan N : ")
	fmt.Scan(&n)

	faktor(n, 1)
}