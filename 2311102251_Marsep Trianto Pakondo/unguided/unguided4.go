package main

import "fmt"

func barisBilanganTurun(n, x int)  {
	if x == 0 {
		return
	} else {
		if x >= 1 {
			fmt.Print(x, " ")
			barisBilanganTurun(n, x-1)
		}
	}
}
func barisBilanganNaik(n, x int)  {
	if x > n  {
		return
	} else {
		if x <= n {
			fmt.Print(x, " ")
			barisBilanganNaik(n, x+1)
		}
	}
}

func main()  {
	var n int

	fmt.Print("Masukkan bilangan N : ")
	fmt.Scan(&n)

	x := n
	barisBilanganTurun(n, x)
	barisBilanganNaik(n, 2)
}