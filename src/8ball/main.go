package main

import (
	m "8ball/lib"
	"fmt"
	h "github.com/skilstak/go/choice"
	c "github.com/skilstak/go/colors"
	i "github.com/skilstak/go/input"
	s "strings"
)

func main() {
	m.PrintWelcome()
	var myArray = []string{
		"Yes",
		"No",
		"Maybe",
		"Certainly",
		"No way!",
		"I don't know...",
	}
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
}
