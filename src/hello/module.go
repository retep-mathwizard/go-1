package hellofunc

import (
	"fmt"
	c "github.com/skilstak/go/colors"
)

func PrintNormal(message string) {
	fmt.Println(message)
}
func FillScreen(message string) {
	for {
		fmt.Print(c.Rc() + message + " ")
	}
}
func PrintForever(message string) {
	for {
		fmt.Println(c.Clear + c.Multi(message))
	}
}
func PrintMulti(message string) {
	fmt.Println(c.Multi(message) + c.X)
}
func PrintRandom(message string) {
	fmt.Println(c.Rc() + message + c.X)
}
