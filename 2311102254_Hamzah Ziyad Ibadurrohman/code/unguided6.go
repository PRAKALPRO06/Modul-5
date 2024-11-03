package main

import "fmt"

func pangkatkan(x, y int) int {
	if y == 0 {
		return 1
	} else {
		return x * pangkatkan(x, y-1)
	}
}
func main() {
	var x, y, hasilnya int
	fmt.Print("Masukkan 2 variabel pangkat: ")
	fmt.Scan(&x, &y)
	hasilnya = pangkatkan(x, y)
	fmt.Print("hasil pangkat:", hasilnya)

}
