//Rakha Yudhistira_2311102010_IF-11-06
package main
import (
	"fmt"
)

func findFactors(N, pembagi int) {

	if pembagi > N {
		return
	}

	if N%pembagi == 0 {
		fmt.Printf("%d ", pembagi)
	}

	findFactors(N, pembagi+1)
}

func main() {
	var N int

	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Printf("Faktor-faktor dari %d adalah: ", N)
	findFactors(N, 1) 
	fmt.Println()
}