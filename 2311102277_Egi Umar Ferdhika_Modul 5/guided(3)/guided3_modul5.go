package main

import "fmt"

func pangkat(n int) int {
	if n == 0 {
		return 1
	}
	return 2 * pangkat(n-1)
}

func main() {
	var n int
	fmt.Print("Masukkab nilai n untuk mencari 2 pangkat n : ")
	fmt.Scanln(&n)
	fmt.Println("Hasil 2 pangkat", n, ":", pangkat(n))
}
