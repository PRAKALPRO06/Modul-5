package main

import "fmt"

func hasilpangkat(x int, y int) int {

	if y == 0 {
		return 1
	}

	return x * hasilpangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x dan y: ")
	fmt.Scan(&x, &y)
	hasil := hasilpangkat(x, y)
	fmt.Printf("%d pangkat %d adalah %d\n", x, y, hasil)
}
