package main

import (
	"fmt"
	"greet"
)

func main() {
	fmt.Println("This is main function")
	fmt.Printf("greet.MORNING= %s", greet.MORNING)
}

func funcApp1() {
	fmt.Println("calling app.funcApp1")
}
