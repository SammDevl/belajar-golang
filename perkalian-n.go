package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&x)
		fmt.Printf("%d ", x*i)
	}
}
