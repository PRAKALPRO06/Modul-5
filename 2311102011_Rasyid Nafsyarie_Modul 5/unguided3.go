//Rasyid Nafsyarie 2311102011 IF 11 06

package main
import (
	"fmt"
)

func findFactors(N, divisor int) {

	if divisor > N {
		return
	}

	if N%divisor == 0 {
		fmt.Printf("%d ", divisor)
	}

	findFactors(N, divisor+1)
}

func main() {
	var N int

	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Printf("Faktor-faktor dari %d adalah: ", N)
	findFactors(N, 1) 
	fmt.Println()
}
