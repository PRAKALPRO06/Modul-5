package main

import "fmt"

//fungsi untuk mencari 2 pangkat
func pangkatDua(n int) int {
	if n == 0 {
		return 1
	}
	return 2 * pangkatDua(n-1)
}

func main() {
	var n int
	fmt.Print("Masukkab nilai n untuk mencari 2 pangkat n : ")
	fmt.Scanln(&n)
	fmt.Println("Hasil 2 pangkat", n, ":", pangkatDua(n))
}
