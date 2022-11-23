package main

import "fmt"

func main() {
	const a = 10
	const b = 30
	//if-else statement
	if a == b {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	//case statement
	switch a {
	case 1:
		fmt.Printf("false")
	case 10:
		fmt.Printf("true")
	default:
		fmt.Printf("hello")
	}
}
