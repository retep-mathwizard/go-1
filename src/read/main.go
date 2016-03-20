package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func getSlice(file string) []string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var slice []string
	err = json.Unmarshal(data, &slice)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return slice
}
func main() {
	slice := getSlice("data.json")
	fmt.Println(slice)
}
