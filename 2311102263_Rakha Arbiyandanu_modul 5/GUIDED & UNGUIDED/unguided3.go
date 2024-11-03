package main
import "fmt"

func cetakFaktor(n, i int){
	if i > n{
		return
	}
	if n%i == 0{
		fmt.Print(i," ")
	}
	cetakFaktor(n, i+1)
}

func main(){
	var n int
	fmt.Print("masukkan bilangan : ")
	fmt.Scan(&n)
	fmt.Print("Faktor dari ", n, ": ")
	cetakFaktor(n, 1)
	fmt.Println()
}