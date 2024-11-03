package main
import (
	"fmt"
)

func printOddNumbers(N, current int) {
	if current > N {
		return
	}
	if current%2 != 0 {
		fmt.Printf("%d ", current)
	}
	printOddNumbers(N, current+1)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Printf("Bilangan ganjil dari 1 hingga %d adalah:\n", N)
	printOddNumbers(N, 1)
	fmt.Println()
}

//Reza Alvonzo 2311102026
