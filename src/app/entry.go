package main

import (
	"fmt"
	"greet"
	"greet/stringutil"
)

func main() {
	fmt.Println("This is main function")
	fmt.Printf("greet.MORNING= %s \n", greet1.MORNING)
	fmt.Printf("Calling external package: %s \n", stringutil.Prefix("calling"))
}

func funcApp1() {
	fmt.Println("calling app.funcApp1")
}
