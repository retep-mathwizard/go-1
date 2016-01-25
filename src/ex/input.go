package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/skilstak/go/input"
	s "strings"
)

func main() {
	input, err := i.Prompt(c.Rc() + "--> " + c.X)
	if err != nil {
		panic(err)
	}
	if s.Contains(input, "hi") {
		fmt.Println(c.Rc() + "Hello!" + c.X)
	} else {
		fmt.Println("")
	}

}
