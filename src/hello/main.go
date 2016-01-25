package main

import (
	"fmt"
	//c  is basically importing it as c then go get
	c "github.com/skilstak/go/colors"
	i "github.com/skilstak/go/input"
	m "hello/lib"
	"os"
	s "strings"
)

func getNameExample(title string) string {
	//first parens are in, second is out. If only one paren then it's only out.
	//don't need := if variable already defined, use =
	name := "Peter"
	//:= infers the type: string, number
	return title + " " + name
}
func printNormal(message string) {
	fmt.Println(message)
}
func fillScreen(message string) {
	for {
		fmt.Print(c.Rc() + message + " ")
	}
}
func printForever(message string) {
	for {
		fmt.Println(c.Clear + c.Multi(message))
	}
}
func printMulti(message string) {
	fmt.Println(c.Multi(message) + c.X)
}
func printRandom(message string) {
	fmt.Println(c.Rc() + message + c.X)
}
func userInput() {
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
func main() {
	name := "World"
	option := ""
	if len(os.Args) > 2 {
		name = os.Args[1]
		option = os.Args[2]
		if s.HasPrefix(option, "-f") {
			printForever("Hello " + name)
		} else if s.HasPrefix(option, "-r") {
			printRandom("Hello " + name)
			fmt.Println("No")
		} else if s.HasPrefix(option, "-m") {
			printMulti("Hello " + name)
		} else if s.HasPrefix(option, "-s") {
			fillScreen("Hello " + name)
		} else {
			printNormal("Hello " + name)
		}
	} else if len(os.Args) > 1 {
		fmt.Println("Add a space then -f,-r,-m,or -s to have more fun!")
		name = os.Args[1]
		printNormal("Hello " + name)
	} else {
		fmt.Println("Add a space and then your name to have more fun!")
		printNormal("Hello " + name)
	}
	m.ImportLine()
	fmt.Println(Statement)
}
