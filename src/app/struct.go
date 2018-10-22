package main

import (
	"fmt"
)

func main() {
	var t1 Test
	t1.test1 = "1"
	employeeOfTheMonth := t1
	employeeOfTheMonth.test1 += " (proactive team player)"
	fmt.Printf("employeeOfTheMonth.test1=%s", employeeOfTheMonth.test1)
	t2 := Test{"1", "3", 2, 4, "5", &employeeOfTheMonth}
	fmt.Printf("AAA %s", t2)
}

type Test struct {
	test1, test3 string
	test2, test4 int
	Test5        string `json:aaaa"`
	myself       *Test
}
