package main
import "fmt"

func main() {
	var n int
	fmt.Print("Masukan N, Untuk pola bintang: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}