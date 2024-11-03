// Meutya Azzahra Efendi
// 2311102166
// IF-11-06
package main

import "fmt"

func printSequence(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	printSequence(n - 1)
	fmt.Print(n, " ")
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scanln(&n)

	fmt.Println("Barisan bilangan:")
	printSequence(n)
}
