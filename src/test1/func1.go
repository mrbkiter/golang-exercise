package main

import (
	"errors"
	"fmt"
	"math"
	"testlib/test2"
)

/* define a circle */
type Circle struct {
	x, y, radius float64
}

/* define a method for circle */
func (circle Circle) area() float64 {
	test2.test2()
	return math.Pi * circle.radius * circle.radius
}

func error1() (int, error) {
	return 1, errors.New("Test error")
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	fmt.Printf("Circle area: %f", circle.area())
	circle2 := Circle{x: 1, y: 1, radius: 5}
	fmt.Println("%f", circle2.area())
	z1 := len("AAAAA")
	fmt.Print(z1)
	x, err := error1()
	if err != nil {
		fmt.Print("%d %s", x, err.Error())
	}
}
