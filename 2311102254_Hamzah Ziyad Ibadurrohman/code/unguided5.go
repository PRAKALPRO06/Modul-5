package main

import "fmt"

func bilanganganjil(n int) {
	if n == 1 {
		fmt.Print(1)
	} else if n%2 == 0 {
		bilanganganjil(n - 1)
	} else {
		bilanganganjil(n - 2)
		fmt.Print(n)
	}

}
func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)
	bilanganganjil(n)

}
