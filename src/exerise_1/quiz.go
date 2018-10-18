package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./quiz.csv")
	if err == nil {
		//fmt.Print(string(data))
		reader := bufio.NewReader(os.Stdin)
		text := string(data)
		lines := strings.Split(text, "\n")
		noOfRightAns := 0
		for index, row := range lines {
			question := strings.Split(row, ",")
			fmt.Printf("Question #%d: %s \n", index+1, question[0])
			userAns, _ := reader.ReadString('\n')
			if strings.Compare(strings.Trim(userAns, " \n"), strings.Trim(question[1], " \n")) == 0 {
				noOfRightAns = noOfRightAns + 1
			}
		}
		fmt.Printf("No of right Answers: %d\n", noOfRightAns)
	}
}
