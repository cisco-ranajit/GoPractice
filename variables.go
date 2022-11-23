package main

import "fmt"

func main() {
	var (
		name   string
		age    int
		weight float64
	)
	fmt.Scanf("%s", &name)
	fmt.Scanf("%d", &age)
	fmt.Scanf("%f", &weight)
	fmt.Println(name)
	fmt.Printf("%d\n", age)
	fmt.Printf("%f\n", weight)
}
