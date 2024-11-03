package main

import "fmt"

func cetakPangkat_226(x, y int) int {
    if y == 0 {
        return 1
    }
    return x * cetakPangkat_226(x, y-1)
}

func main() {
    var x, y int
    
    fmt.Print("Masukkan dua bilangan (x y): ")
    fmt.Scan(&x, &y)
    
    result := cetakPangkat_226(x, y)
    
    fmt.Printf("Hasil %d dipangkatkan %d adalah: %d\n", x, y, result)
}