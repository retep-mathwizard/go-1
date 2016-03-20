package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	info, err := ioutil.ReadFile("data")
	if err != nil {
		panic(err)
	}
	s := string(info)
	fmt.Println(s)
}
