package main

import (
	"fmt"
	"strings"
)

type human interface {
	Read(p string) string
}

//Professor professor
type Professor struct {
	name string
}

func (Professor) Read(p string) string {
	fmt.Printf("Professor said %s\n", p)
	return "Professor"
}

//Read read a message and then export value
func Read(p string) string {
	fmt.Printf("Message %s\n", p)
	return strings.Replace(p, "a", "a1", -1)
}

func test1(h human) {
	h.Read("Hello world")
	h1 := h.(Professor)
	fmt.Printf("This is unbox to Professor %s\n", h1.name)
}

func main() {
	p := Professor{name: "1"}
	test1(p)
	p1 := &p
	var p2 *Professor
	p2 = &p
	fmt.Printf("p.name=%s, p1.name=%s,p2.name=%s\n", p.name, p1.name, p2.name)
}
