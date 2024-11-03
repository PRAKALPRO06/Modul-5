package main 

  

import "fmt" 

  

// Fungsi Utama 

func main() { 

    var x, y int 

    fmt.Print("Muhammad Djoko Susilo/2311102212 \n") 

    fmt.Print("Masukkan bilangan x: ") 

    fmt.Scan(&x) 

    fmt.Print("Masukkan bilangan y: ") 

    fmt.Scan(&y) 

    hasil := pangkat(x, y) 

    fmt.Printf("%d^%d = %d\n", x, y, hasil) 

} 

  

// Fungsi rekursif untuk menghitung x pangkat y 

func pangkat(x, y int) int { 

    if y == 0 { 

        return 1 // Basis rekursi: x^0 = 1 

    } else { 

        return x * pangkat(x, y-1) // Menghitung x * x^(y-1) 

    } 

} 

 