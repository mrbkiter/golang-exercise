package main

import "fmt"

func main() {
	const HELLO_WORLD = "Hello World"

	var x = 20.0
	var xx int = 2

	var y = 42
	z := 12
	fmt.Println(x, z)
	fmt.Println(xx)
	fmt.Println(y)
	fmt.Printf("x is of type %f %T\n", x, x)
	fmt.Printf("y is of type %T\n", y)

	switch x {
	case 10:
		fmt.Print("This is 10")
	case 20:
		fmt.Print("Bingo")

	}
}
