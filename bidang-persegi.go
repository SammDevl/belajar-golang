package main
import "fmt"

func main() {
	var n, sisi, keliling, luas int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&sisi)
		keliling = 4 * sisi
		luas = sisi * sisi
		fmt.Println(luas, keliling)
	}
	
}	