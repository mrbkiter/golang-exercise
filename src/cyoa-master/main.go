package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, error := os.Open(".\\src\\cyoa-master\\gopher.json")
	if error != nil {
		panic(error)
	}
	bcontent, error := ioutil.ReadAll(file)
	guideMap := make(map[string]Story)
	error1 := json.Unmarshal(bcontent, &guideMap)
	if error1 != nil {
		fmt.Printf(error1.Error())
	} else {
		fmt.Printf("%v\n", guideMap)
	}

}

//Story of a book.
type Story struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

//Option that help contain bookmark
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
