// 2311102012
package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghasilkan barisan angka
func buatBarisan(angkaAwal, angkaSekarang int) []int {
	// Jika sudah mencapai angka 1, mulai membangun barisan balik
	if angkaSekarang == 1 {
		return []int{1}
	}
	// Panggil fungsi secara rekursif untuk menghasilkan barisan angka
	barisan := append([]int{angkaSekarang}, buatBarisan(angkaAwal, angkaSekarang-1)...)
	return append(barisan, angkaSekarang)
}

// Fungsi pembantu untuk mengubah slice angka menjadi string
func ubahBarisanKeString(barisan []int) string {
	hasil := ""
	for _, angka := range barisan {
		hasil += fmt.Sprintf("%d ", angka)
	}
	return hasil
}

func main() {
	var angka int
	// Meminta input dari pengguna
	fmt.Print("Masukkan sebuah angka positif: ")
	fmt.Scanln(&angka)

	// Memeriksa jika input valid
	if angka <= 0 {
		fmt.Println("Masukkan angka positif yang lebih besar dari 0.")
		return
	}

	// Memanggil fungsi dan menampilkan hasil
	fmt.Println("Hasil:", ubahBarisanKeString(buatBarisan(angka, angka)))
}
