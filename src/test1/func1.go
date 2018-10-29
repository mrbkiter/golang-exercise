package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

/* define a circle */
type Circle struct {
	x, y, radius float64
}

/* define a method for circle */
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func error1() (int, error) {
	return 1, errors.New("Test error")
}

func main() {
	domain1 := "github.com"
	domain2 := "test.github.com"
	if len(domain2) > len(domain1) && strings.Compare(domain2[len(domain2)-(len(domain1)+1):], "."+domain1) == 0 {
		fmt.Print("TRUE")
	} else {
		fmt.Print("FALSE")
	}
}
func main1() {
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
