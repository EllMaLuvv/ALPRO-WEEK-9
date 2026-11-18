package main
import "fmt"

func main () {
var x, y int
var xfaky, yfakx bool

fmt.Scan(&x, &y)
if y%x == 0 {
	xfaky = true
}
if x%y == 0 {
	yfakx = true
}
fmt.Println(xfaky)
fmt.Println(yfakx)
}