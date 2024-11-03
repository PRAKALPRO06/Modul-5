package main
import (
	"fmt"
)

func printSeries(N, current int) {
	if current == 1 {
		fmt.Printf("%d ", current)
		return
	}

	fmt.Printf("%d ", current)
	printSeries(N, current-1)
	fmt.Printf("%d ", current)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Printf("Barisan bilangan dari %d hingga 1 dan kembali ke %d:\n", N, N)
	printSeries(N, N)
	fmt.Println()
}
//Reza Alvonzo 2311102026
