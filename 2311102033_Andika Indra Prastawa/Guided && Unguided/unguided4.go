package main

import "fmt"

func printDescending(n int) {
	if n < 1 {
		return
	}
	fmt.Print(n, " ")
	printDescending(n - 1)
}

func printAscending(n, current int) {
	if current > n {
		return
	}
	fmt.Print(current, " ")
	printAscending(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&n)

	fmt.Println("Barisan bilangan dari", n, "hingga 1 dan kembali ke", n, ":")
	printDescending(n)
	printAscending(n, 1)
	fmt.Println()
}
