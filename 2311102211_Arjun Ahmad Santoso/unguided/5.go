package main 

import "fmt"

func cetakGanjil(n int) {
	if n < 0 {
		return
	}
	cetakGanjil(n-2)
	fmt.Print(n, " ")
}

func main() {
	var n int
	fmt.Print("n: ")
	fmt.Scan(&n)
	if n % 2 == 0 {
		n = n - 1
	}
	cetakGanjil(n)
}