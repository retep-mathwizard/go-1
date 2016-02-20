package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type car struct {
	Name   string   `json:"name"`
	Wheels int      `json:"wheels"`
	Awards []string `json:"awards"`
}

func main() {
	// read the file, using the ioutil pacakge  https://golang.org/pkg/io/ioutil/#ReadFile
	bytes, err := ioutil.ReadFile("file.json")
	if err != nil {
		fmt.Println("There was an error reading the file, make sure it is in the same directory as main.go, or you have the correct path stated")
		fmt.Println(err)
		panic(err)
	}
	honda := &car{}
	err = json.Unmarshal(bytes, honda)
	//decodes it
	if err != nil {
		// error unmarshalling the json, do something
		fmt.Println("There was an error decoding the file: make sure you have a correct struct formatting, and uppercase letters")
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(honda.Name)
	fmt.Println(honda.Wheels)
	fmt.Println(honda.Awards)
}
