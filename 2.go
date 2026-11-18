package main
import "fmt"

func main () {
	var not, gen string
	var n int
fmt.Scan(&n)
if n > 0 || n % 2 != 0 {
	not = "Bukan"
	fmt.Println(not)
}
if n % 2 == 0 && n < 0 {
	gen = "Genap Negatif"
	fmt.Println(gen)
}
}