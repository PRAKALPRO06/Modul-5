package main
import "fmt"

func printBarisanBilanganGanjil_dari_1_hingga_N(n, i int) {
	if i > n {
		return
	}
	if i%2 != 0 {
		fmt.Print(i, " ")
	}
	printBarisanBilanganGanjil_dari_1_hingga_N(n, i+1)
}

func main() {
	var n int
	fmt.Print("Input N, untuk bilangan bulat positif nya: ")
	fmt.Scanln(&n)
	fmt.Print("Keluaran : ")
	printBarisanBilanganGanjil_dari_1_hingga_N(n, 1)
	fmt.Println()
}