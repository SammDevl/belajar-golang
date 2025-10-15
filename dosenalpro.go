package main
import "fmt"

func main() {
var x, y float64
	fmt.Scan(&x, &y)

	f := (22 / y) / (7 + 2*x) + (x * y)

	fmt.Println("Tim Dosen AlPro_1", f)
}