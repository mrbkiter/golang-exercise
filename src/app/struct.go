package main

import (
	"fmt"
)

func main() {
	t1 := Test{test1: "test1"}
	pt1 := t1
	t1.test1 = "2222"
	fmt.Printf("t1.test1 = %s, pt1=%v\n", t1.test1, pt1.test1)

	x := 1.5
	square(&x)

	fmt.Printf("x = %f\n", x)

	x1, y1 := 2, 3
	fmt.Printf("&x1=%p, &y1=%p\n", &x1, &y1)
	swap(&x1, &y1)
	fmt.Printf("x1=%d, y1=%d\n", x1, y1)

	fmt.Printf("-------------------\n")
	test1 := Test{test1: "test 1", test2: 0}
	test1.assign1()
	fmt.Printf("assign1 test1.test2=%d\n", test1.test2)
	test1.assign1Pointer()
	fmt.Printf("assign1Pointer test1.test2=%d\n", test1.test2)
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x *int, y *int) {
	fmt.Printf("In function: &x1=%p, &y1=%p\n", x, y)
	z := *x
	*x = *y
	*y = z
}

type Test struct {
	test1 string
	test2 int
}

func (v Test) assign1() {
	v.test2 = 1
}

func (v *Test) assign1Pointer() {
	v.test2 = 1
}
