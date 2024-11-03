package main
import "fmt"

func printBarisanBilanganGanjil(n int, bilangan int){
	if bilangan > n{
		return
	}

	fmt.Print(bilangan, " ")
	printBarisanBilanganGanjil(n, bilangan+2)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif  : ")
	fmt.Scan(&N)
	printBarisanBilanganGanjil(N, 1)
	fmt.Println()
}