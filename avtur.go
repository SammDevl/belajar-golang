package main
import "fmt"

func main() {
	var pound float64
	fmt.Scan(&pound)

	kilogram := pound * 0.45359237
	liter := kilogram / 0.80

	fmt.Printf("%.6f kg, %.6f l\n", kilogram, liter)
}