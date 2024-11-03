//2311102037 BRIAN FARREL EVANDHIKA IF 11 06
package main

import (
    "fmt"
)

// Fungsi rekursif untuk mencetak satu baris bintang
func printStars(n int) {
    if n == 0 {
        return
    }
    fmt.Print("*")
    printStars(n - 1)
}

// Fungsi rekursif untuk mencetak pola bintang bertingkat
func printPattern(n, current int) {
    if current > n {
        return
    }
    // Cetak bintang sejumlah nilai 'current'
    printStars(current)
    fmt.Println() // Pindah ke baris berikutnya
    // Rekursi untuk baris berikutnya
    printPattern(n, current + 1)
}

func main() {
    var n int
    fmt.Print("Masukkan nilai N: ")
    fmt.Scan(&n)
    printPattern(n, 1)
}
