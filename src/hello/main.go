package main

import (
	"fmt"
	//c  is basically importing it as c then go get
	m "hello/lib"
	"os"
	s "strings"
)

func main() {
	name := "World"
	option := ""
	if len(os.Args) > 2 {
		name = os.Args[1]
		option = os.Args[2]
		if s.HasPrefix(option, "-f") {
			m.PrintForever("Hello " + name)
		} else if s.HasPrefix(option, "-r") {
			m.PrintRandom("Hello " + name)
			fmt.Println("No")
		} else if s.HasPrefix(option, "-m") {
			m.PrintMulti("Hello " + name)
		} else if s.HasPrefix(option, "-s") {
			m.FillScreen("Hello " + name)
		} else {
			m.PrintNormal("Hello " + name)
		}
	} else if len(os.Args) > 1 {
		name = os.Args[1]
		m.PrintNormal("Hello " + name)
	} else {
		m.PrintNormal("Hello " + name)
	}
}
