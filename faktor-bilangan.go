package main

import "fmt"

func main() {
	var n int
	var cek bool
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		cek = n % i == 0
		fmt.Println(i, cek)
	}
}