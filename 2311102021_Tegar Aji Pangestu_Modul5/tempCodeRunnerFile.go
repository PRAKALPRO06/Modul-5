package main
import "fmt"

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	} else if y == 1 {
		return x
	} else {
		return x * pangkat(x, y-1)
	}
}

func main() {
	var x, y int
	fmt.Print("Masukkan x: ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan y: ")
	fmt.Scanln(&y)

	fmt.Println(pangkat(x, y))
}