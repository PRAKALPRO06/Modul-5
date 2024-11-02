package main 

import "fmt"

func pangkat(a, b int) int {
	if b == 1 {
		return a
	}
	result := a * pangkat(a, b-1)
	return result
}

func main() {
	var a, b int
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Print("b: ")
	fmt.Scan(&b)
	result := pangkat(a, b)
	fmt.Print("Hasil dari ", a, " pangkat ", b, " adalah ", result)
}