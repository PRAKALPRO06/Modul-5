package main
import "fmt"

func FungsiPangkat_y(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * FungsiPangkat_y(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan x dan y Sekaligus, Spasi jejeran: ")
	fmt.Scanf("%d %d", &x, &y)
	fmt.Printf(" %d Pangkat %d Hasilnya: %d\n", x, y, FungsiPangkat_y(x, y))
}