////Wisnu Rananta Raditya Putra (2311102013) IF-11-06

package main
import "fmt"

func cetakBintang(n, barisDiCetak_2311102013 int) {
	if barisDiCetak_2311102013 > n {
		return
	}
	for i := 0; i < barisDiCetak_2311102013; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	cetakBintang(n, barisDiCetak_2311102013+1)
}
func main() {
	var n int
	fmt.Print("Masukkan jumlah baris :")
	fmt.Scan(&n)
	cetakBintang(n, 1)
}
