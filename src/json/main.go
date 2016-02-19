package main

import (
	"fmt"
	"io/ioutil"
	//file reading package
	"encoding/json"
	//.json decoding package
)

type MyStruct struct {
	X    int    `json:"X"`
	Y    int    `json:"Y"`
	Name string `json:"Name"`
	//tells it that it's a json type, and should look for "X" in the json file
}

func main() {
	file, err := ioutil.ReadFile("file.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(file)
	//not decoded
}
