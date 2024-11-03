package main 

  

import "fmt" 

  

// Procedure untuk mencetak pola bintang secara rekursif 

func pola(i, N int, s string) { 

    if i <= N { 

        s = "*" + s 

        fmt.Println(s) 

        pola(i+1, N, s) 

    } 

} 

  

func main() { 

    var N int 

    fmt.Print("Muhammad Djoko Susilo/2311102212 \n") 

    fmt.Print("Masukkan jumlah baris N untuk pola bintang: ") 

    fmt.Scanln(&N) 

     

    fmt.Println("Hasil Pola Bintang:") 

    pola(1, N, "") 

} 

 