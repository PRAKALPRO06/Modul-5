package main 

import "fmt"

func cetakPolaNaik(n int) {
	if n == 0 {
		return
	}
	cetakPolaNaik(n-1)
	fmt.Print(n, " ")
}

func cetakPolaTurun(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	cetakPolaTurun(n-1)
}

func main() {
	var n int
	fmt.Print("n: ")
	fmt.Scan(&n)
	
	cetakPolaTurun(n)
	cetakPolaNaik(n)
}