package main 

import "fmt"

func cetakFaktor(n *int, i int) {
	if i > *n {
		return
	}
	if *n % i == 0 {
		fmt.Print(i, " ")
	}
	cetakFaktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("n: ")
	fmt.Scan(&n)
	fmt.Println("Faktor dari", n, ":")
	cetakFaktor(&n, 1)
}