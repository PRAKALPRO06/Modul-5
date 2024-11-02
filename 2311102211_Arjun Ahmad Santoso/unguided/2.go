package main

import "fmt"

func cetakPolaBintang(n int) {
	if n == 0 { return }
	cetakPolaBintang(n-1)
	for i:=0; i<n; i++ {
		fmt.Print("*")
	}
	fmt.Print("\n")
}

func main() {
	var n int
	fmt.Print("n: ")
	fmt.Scanln(&n)
	cetakPolaBintang(n)
}