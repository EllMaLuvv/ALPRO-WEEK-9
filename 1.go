package main
import "fmt"

var orang, motor int
func main () {
	fmt.Scan(&orang)
	motor = orang / 2
	if orang % 2 != 0 {
		motor = motor + 1
	}
	fmt.Println(motor)
}
