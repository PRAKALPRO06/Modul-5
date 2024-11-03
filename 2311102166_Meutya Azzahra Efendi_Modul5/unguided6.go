// Meutya Azzahra Efendi
// 2311102166
// IF-11-06
package main

import "fmt"

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int

	fmt.Println("Masukkan bilangan dasar (x):")
	fmt.Scanln(&x)
	fmt.Println("Masukkan pangkat (y):")
	fmt.Scanln(&y)

	hasil := pangkat(x, y)
	fmt.Printf("%d pangkat %d adalah: %d\n", x, y, hasil)
}
