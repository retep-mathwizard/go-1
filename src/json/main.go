package main

import (
	"fmt"
	"io/ioutil"
	//file reading package
	//"encoding/json"
	//"io"
	"os"
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
	//file2 := io.Reader(file)
	//fmt.Println(file2)
	//error
	file4, _ := os.Open("file.json")
	fmt.Println(file4)
}

//returned
//file: [123 10 32 32 88 58 53 44 10 32 32 89 58 49 48 44 10 32 32 78 97 109 101 58 34 98 111 98 34 10 125 10]
//file4: &{0xc820010390}
//
//
//
