package main

import "fmt"

func printNtoN(n, current int) {
	fmt.Printf("%d ", current)
	if current == 1 {
		return
	}
	printNtoN(n, current-1)
	fmt.Printf("%d ", current)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)

	fmt.Print("Keluaran: ")
	printNtoN(n, n)
	fmt.Println()
}
