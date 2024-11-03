package main
import "fmt"

func deretBilangan(n, bilangan int){
	if bilangan > n {
		return
	}

	fmt.Print(bilangan, " ")
	if bilangan == 1{
		balikPola(2, n)
	}else {
		deretBilangan(n, bilangan-1)
	}
}

func balikPola(bilangan, n int){
	if bilangan > n {
		return
	}

	fmt.Print(bilangan, " ")
	balikPola(bilangan+1, n)
}

func main(){
	var n int
	fmt.Print("Masukkan angka : ")
	fmt.Scanln(&n)

	deretBilangan(n, n)
	fmt.Println()
}