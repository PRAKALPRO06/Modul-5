package main  

  

import "fmt" 

  

//fungsi untuk menghitung penjumlahan 1 hingga n 

func jumlahRekursif(n int) int { 

    if n == 1 { 

        return 1 

    } 

    return n + jumlahRekursif(n-1) 

} 

  

func main(){ 

    var n int 

    fmt.Print("Muhammad Djoko Susilo/2311102212 \n") 

    fmt.Print("Masukan nilai n untuk penjumlahan 1 hingga n: ") 

    fmt.Scanln(&n) 

    fmt.Println("Hasil Penjumlahan: ", jumlahRekursif(n)) 

} 