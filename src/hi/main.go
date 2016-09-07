package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	"github.com/skilstak/go-input"
)

func main() {
	name := input.Ask(c.B3 + "What is your name? > ")
	fmt.Println("Nice to meet you " + name + "!!" + c.X)
}
