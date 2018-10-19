package greet1

import (
	"fmt"
)

//MORNING is exposed variable
var MORNING = "Morning is exposed"
var moring = "private variable"

func helloWorld() {
	fmt.Println("Hello world from greet1.go")
}
