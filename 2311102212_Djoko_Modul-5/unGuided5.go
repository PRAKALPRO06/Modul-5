package main 

  

import "fmt" 

  

// Fungsi Utama 

func main() { 

    var n int 

    fmt.Print("Muhammad Djoko Susilo/2311102212 \n") 

    fmt.Print("Masukkan bilangan ganjil ke-n: ") 

    fmt.Scan(&n) 

    ganjil(n) 

} 

  

// Prosedur rekursif untuk mencetak bilangan ganjil dari 1 hingga ganjil ke-n 

func ganjil(n int) { 

    if n == 1 { 

        fmt.Print(1, " ")  // Basis rekursi, bilangan ganjil pertama adalah 1 

    } else { 

        if n%2 == 0 { 

            ganjil(n - 1)  // Jika n genap, abaikan dan panggil ganjil(n-1) 

        } else { 

            ganjil(n - 2)  // Panggil ganjil(n-2) untuk mencetak bilangan ganjil sebelumnya 

            fmt.Print(n, " ") 

        } 

    } 

} 

 