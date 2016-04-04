package info

import (
	"fmt"
	i "github.com/retep-mathwizard/utils/input"
	c "github.com/skilstak/go/colors"
	"reflect"
)

//%s = colors, %f = strings
func GetName() (string, string) {
	p1name := i.StringInput(c.CL + "Player 1, what is your name? > ")
	fmt.Println(c.CL)
	p2name := i.StringInput("Player 2, what is your name? > ")
	fmt.Println(c.CL)
	return p1name, p2name
}

func PrintType(item interface{}) {
	Objecttype := reflect.TypeOf(item)
	fmt.Println(Objecttype)
}
