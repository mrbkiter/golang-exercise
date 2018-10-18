package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	timer1 := time.NewTimer(30 * time.Second)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press any key to start")
	reader.ReadByte()
	data, err := ioutil.ReadFile("./quiz.csv")

	answerCh := make(chan string)
	if err == nil {
		//fmt.Print(string(data))
		reader := bufio.NewReader(os.Stdin)
		text := string(data)
		lines := strings.Split(text, "\n")
		noOfRightAns := 0
	problemloop:
		for index, row := range lines {
			question := strings.Split(row, ",")
			go func() {
				fmt.Printf("Question #%d: %s \n", index+1, question[0])
				userAns, _ := reader.ReadString('\n')
				answerCh <- userAns
			}()

			select {
			case <-timer1.C:
				fmt.Println("Time is over")
				break problemloop
			case answer := <-answerCh:
				if strings.Compare(strings.Trim(answer, " \n"), strings.Trim(question[1], " \n")) == 0 {
					noOfRightAns = noOfRightAns + 1
				}
			}
		}

		fmt.Printf("No of right Answers: %d\n", noOfRightAns)
	}

	fmt.Println("Ended")
}
