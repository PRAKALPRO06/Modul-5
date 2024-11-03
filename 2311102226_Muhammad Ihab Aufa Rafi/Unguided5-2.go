package main

import (
	"fmt"
)

func cetakBintang_226(n int) string {
    if n == 0 {
        return ""
    }
    return "*" + cetakBintang_226(n-1)
}

func printPattern(baris, totalBaris int) {
    if baris > totalBaris {
        return
    }
    fmt.Println(cetakBintang_226(baris))    
    printPattern(baris+1, totalBaris)
}

func main() {
    var n int
    
    for {
        fmt.Print("Masukkan jumlah baris (n): ")
        fmt.Scan(&n)
        
        if n > 0 {
            break
        }
        fmt.Println("Masukkan harus lebih besar dari 0!")
    }
    
    fmt.Printf("\nPola Bintang untuk n = %d:\n", n)
    printPattern(1, n)
}