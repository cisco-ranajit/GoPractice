package main

import "fmt"

func main() {
	var (
		a, b, c int
	)
	fmt.Println("enter number for a")
	fmt.Scanf("%d\n", &a)
	fmt.Println("enter number for b")
	fmt.Scanf("%d\n", &b)

	c = a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	c = a - b
	fmt.Printf("%d - %d = %d\n", a, b, c)
	c = a / b
	fmt.Printf("%d / %d = %d\n", a, b, c)
	c = a % b
	//fmt.Printf("%d % %d = %d", a, b, c)
	c = a * b
	fmt.Printf("%d * %d = %d", a, b, c)
}
