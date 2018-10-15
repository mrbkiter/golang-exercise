package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	timer1 := time.NewTimer(10 * time.Second)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press any key to start")
	reader.ReadByte()
	go func() {
		<-timer1.C
		fmt.Println("End timer")
	}()
	fmt.Println("Ended")
}
