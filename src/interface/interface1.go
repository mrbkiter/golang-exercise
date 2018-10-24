package interfacep

import (
	"fmt"
	"strings"
)

type human interface {
	Read(p string) string
}

//Professor professor
type Professor struct {
}

func (d Professor) Read(p string) string {
	fmt.Printf("Professor said %s\n", p)
	return "Professor"
}

//Read read a message and then export value
func Read(p string) string {
	fmt.Printf("Message %s\n", p)
	return strings.Replace(p, "a", "a1", -1)
}
