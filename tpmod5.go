package main

import "fmt"

func main1() {
	var N int
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		fmt.Printf("%d\n", x * y)
	}
}

func main2() {
	var N int
	fmt.Scan(&N)

	fmt.Print(N) 
	for i := N - 1; i >= 1; i-- {
		fmt.Printf(" x %d", i)
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	for i := N - 1; i > 1; i-- {
		N = N * i
	}
	fmt.Print(N)
}