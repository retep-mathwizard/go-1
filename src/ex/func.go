package main

import (
	"fmt"
)

func function(title string) string {
	//first parens are what is required, second is what type is returned.
	//If only one paren then it's only out.
	//don't need := if variable already defined, use =
	name := "Peter"
	//:= infers the type: string, number, etc
	return title + " " + name
}
func main() {
	fmt.Println(function("Hello"))
}
