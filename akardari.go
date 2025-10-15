package main
import "fmt"

func main() {
    var a, b int
	fmt.Scan(&a, &b)

	if a*a == b {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}