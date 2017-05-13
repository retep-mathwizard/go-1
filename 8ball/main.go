package main

import (
	//m "8ball/lib"
	"fmt"
	//h "github.com/skilstak/go/choice"
	//c "github.com/skilstak/go/colors"
	//i "github.com/skilstak/go/input"
	//s "strings"
	"encoding/json"
	"io/ioutil"
)

type easterEgg struct {
	keys     []string `json:"keys"`
	response string   `json:"response"`
}
type eggs struct {
	answers []easterEgg `json:"answers"`
}

func main() {
	bytes, err := ioutil.ReadFile("data/answers.json")
	if err != nil {
		fmt.Println("There was an error reading the file")
		fmt.Println("make sure it is in the same directory as main.go, or you have the correct path stated")
		panic(err)
	}
	jsonClass := &eggs{}
	err = json.Unmarshal(bytes, jsonClass)
	//decodes it
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			fmt.Printf("%v: %s <<--ERROR %s\n", e, bytes[:e.Offset], bytes[e.Offset:])
		} else {
			fmt.Println("There wan error decoding the file")
			// error unmarshalling the json, do something
		}
		panic(err)
	}
	fmt.Println(jsonClass.answers)
	/*
		var myArray = []string{
			"Yes",
			"No",
			"Maybe",
			"Certainly",
			"No way!",
			"I don't know...",
		}
		for _, easterdict := range secrets {
			for _, element {
				dict := secrets[easterdict]
				for _ := range dict {

				for i := range result {
					fmt.Println(result[i])
				}
				dict[element]


		result := s.Split(, "|")




		for {
			question, err := i.Prompt(c.Rc() + "--> ")
			if err != nil {
				panic(err)
			} else {
				if s.Contains(question, "die") {
					fmt.Println("let's not talk about death...")
				} else {
					randomItem := h.Choice(myArray)
					fmt.Println(c.Rc() + c.CL + randomItem + c.X)
				}
			}
		}
	*/
}
