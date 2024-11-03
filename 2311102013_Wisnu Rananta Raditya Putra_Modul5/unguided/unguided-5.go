//Wisnu Rananta Raditya Putra (2311102013) IF-11-06

package main
import "fmt"

func printBarisanBilanganGanjil(n int, bilangan_2311102013 int) {
	if bilangan_2311102013 > n {
		return
	}
	fmt.Print(bilangan_2311102013, " ")
	printBarisanBilanganGanjil(n, bilangan_2311102013+2)
}
func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif : ")
	fmt.Scan(&N)
	printBarisanBilanganGanjil(N, 1)
	fmt.Println()
}
