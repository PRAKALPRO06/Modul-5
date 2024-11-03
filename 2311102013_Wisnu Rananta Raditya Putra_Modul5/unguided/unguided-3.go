//Wisnu Rananta Raditya Putra (2311102013) IF-11-06

package main
import "fmt"

func cetakFaktor(n_2311102013, i int) {
	if i > n_2311102013 {
		return
	}
	if n_2311102013%i == 0 {
		fmt.Print(i, " ")
	}
	cetakFaktor(n_2311102013, i+1)
}
func main() {
	var n int
	fmt.Print("masukkan bilangan : ")
	fmt.Scan(&n)
	fmt.Print("Faktor dari ", n, ": ")
	cetakFaktor(n, 1)
	fmt.Println()
}
