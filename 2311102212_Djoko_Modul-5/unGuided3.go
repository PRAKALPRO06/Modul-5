package main 

  

import "fmt" 

  

// Procedure untuk mencari dan mencetak faktor dari bilangan n 

func faktor(n, i int) { 

    if i <= n { 

        if n%i == 0 { 

            fmt.Println(i) 

        } 

        faktor(n, i+1) 

    } 

} 

  

func main() { 

    var n int 

    fmt.Print("Muhammad Djoko Susilo/2311102212 \n") 

    fmt.Print("Masukkan nilai n untuk mencari faktor-faktornya: ") 

    fmt.Scanln(&n) 

     

    fmt.Println("Faktor-faktor dari", n, "adalah:") 

    faktor(n, 1) 

} 

 