package main
import "fmt"

func FaktorBilanganDari_N(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	FaktorBilanganDari_N(n, i+1)
}

func main() {
	var n int
	fmt.Print("Input bilangan bulat positif N: ")
	fmt.Scanln(&n)
	fmt.Print("Bilangan yang jadi faktor dari N = ", n )
	fmt.Println()

	fmt.Print("Yaitu : ")
	FaktorBilanganDari_N(n, 1)
	fmt.Println()
}