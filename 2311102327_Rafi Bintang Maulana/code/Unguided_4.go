package main
import "fmt"

func cetakBilanganDari_N_hingga_1_danKembaliKe_N(n int) {
	if n <= 0 {
		return
	}
	fmt.Print(n, " ")
	cetakBilanganDari_N_hingga_1_danKembaliKe_N(n - 1)
	if n > 1 {
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Print("Input N, bilangan bulat yang positif : ")
	fmt.Scanln(&n)
	fmt.Print("Keluaran: ")
	cetakBilanganDari_N_hingga_1_danKembaliKe_N(n)
	fmt.Println()
}