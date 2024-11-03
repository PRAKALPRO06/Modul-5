//Wisnu Rananta Raditya Putra (2311102013) IF-11-06

package main
import "fmt"

func hasilpangkat(x_2311102013 int, y int) int {
	if y == 0 {
		return 1
	}
	return x_2311102013 * hasilpangkat(x_2311102013, y-1)
}
func main() {
	var x, y int
	fmt.Print("Masukkan nilai x dan y: ")
	fmt.Scan(&x, &y)
	hasil := hasilpangkat(x, y)
	fmt.Printf("%d pangkat %d adalah %d\n", x, y, hasil)
}
