package main

import (
	"fmt"
	"galc"
	c "github.com/skilstak/go/colors"
	i "github.com/skilstak/go/input"
	"strconv"
	s "strings"
)

func main() {
	fmt.Println("[a] addition")
	fmt.Println("[s] subtraction")
	fmt.Println("[m] multiplication")
	fmt.Println("[d] division")
	option, err := i.Prompt(c.Rc() + "--> " + c.X)
	if err != nil {
		panic(err)
	}
	strnumber1, err := i.Prompt(c.Rc() + "First number > " + c.X)
	if err != nil {
		panic(err)
	}
	strnumber2, err := i.Prompt(c.Rc() + "Second number > " + c.X)
	if err != nil {
		panic(err)
	}
	floatnumber1, err := strconv.ParseFloat(strnumber1, 64)
	floatnumber2, err := strconv.ParseFloat(strnumber2, 64)
	if s.Contains(option, "a") {
		galc.Add(floatnumber1, floatnumber2)
	}
	if s.Contains(option, "m") {
		galc.Multiply(floatnumber1, floatnumber2)
	}
	if s.Contains(option, "d") {
		galc.Divide(floatnumber1, floatnumber2)
	}
	if s.Contains(option, "s") {
		galc.Subtract(floatnumber1, floatnumber2)
	}
}
