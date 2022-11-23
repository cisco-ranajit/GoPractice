package main

import "fmt"

func main() {
	var (
		a, b, c int
	)
	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%d\n", &b)
	fmt.Scanf("%d\n", &c)
	avg := (a + b + c) / 3

	fmt.Printf("%d", avg)
}
