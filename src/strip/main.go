package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello-there"
	for index, item := range str {
		if item == rune('-') || item == rune('_') {
			value := string(str[index+1])
			newvalue := strings.ToUpper(value)
			str = str[:index] + newvalue + str[index+2:]
		}
	}
	fmt.Println(str)
}
