// Meutya Azzahra Efendi
// 2311102166
// IF-11-06
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&n)

	fmt.Println("Barisan bilangan ganjil:")
	cetakGanjil(1, n)
}

func cetakGanjil(i int, n int) {
	if i > n {
		return
	}
	fmt.Print(i, " ")
	cetakGanjil(i+2, n)
}
