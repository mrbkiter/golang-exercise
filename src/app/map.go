package main

import (
	"fmt"
)

func main2() {
	ages := make(map[string]int)
	ages["albert"] = 12
	ages["vu"] = 14
	ages["vu"]++
	fmt.Printf("age[vu] = %d\n", ages["vu"])
	/* delete(ages, "vu")
	fmt.Printf("age[vu] = %d", ages["vu"]) */
	for k, xv := range ages {
		fmt.Printf("key = %s, value = %d\n", k, xv)
	}
	v, ok := ages["vu1"]
	if ok {
		fmt.Printf("%d", v)
	} else {
		fmt.Println("Not found")
	}

}
