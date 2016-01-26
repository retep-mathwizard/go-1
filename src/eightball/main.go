package main

import (
	"fmt"
	h "github.com/skilstak/go/choice"
	c "github.com/skilstak/go/colors"
)

func main() {
	welcome := `Hello! Welcome to the magic 8Ball!`
	fmt.Println(welcome)
	var myArray = []string{
		"Yes",
		"No",
		"Maybe",
		"Certainly",
		"No way!",
		"I don't know...",
	}
	item := h.Choice(myArray)
	fmt.Println(c.Rc() + item + c.X)

}
