package strings

import (
	"strings"
)

//ContainEnd compare 2 string input1 and input2. If input1 contains input2 at End of its position, return true. Otherwise return false
func ContainEnd(input1 string, input2 string) bool {
	if len(input1) > len(input2) && strings.Compare(input1[len(input1)-(len(input2)):], input2) == 0 {
		return true
	} else {
		return false
	}
}
