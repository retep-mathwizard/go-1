package works

import (
	"fmt"
)

func printType(x interface{}) bool {
	stringX := false
	switch v := x.(type) {
	case int:
		fmt.Println("int data", v)
	case string:
		stringX = true
		fmt.Println("string data", v)
	default:
		fmt.Println("I don't like your type", v)
	}
	return stringX
}
func main() {
	var x string = "hello"
	printType(x)
	if isString == true {
		fmt.Println("It's a string!")
	}
}
