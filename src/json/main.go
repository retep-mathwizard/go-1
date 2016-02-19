package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type car struct {
	Name   string   `json:"name"`
	Wheels int      `json"wheels"`
	Awards []string `json:"awards"`
}

func main() {
	// read the file, using the ioutil pacakge  https://golang.org/pkg/io/ioutil/#ReadFile
	bytes, err := ioutil.ReadFile("file.json")
	if err != nil {
		fmt.Println(err)
		panic(err)
		// failed to read file, do something with err
	}
	honda := &car{}
	err = json.Unmarshal(bytes, honda)
	if err != nil {
		// error unmarshalling the json, do something
	}
	fmt.Println(honda.Name)
	fmt.Println(honda.Wheels)
	fmt.Println(honda.Awards)
}
