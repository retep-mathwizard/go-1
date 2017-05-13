package main

import (
	"fmt"
	"strings"
)

func rot(number int, str string) string {
	caps := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower := "abcdefghijklmnopqrstuvwxyz"
	if number < 0 {
		number = 26 - number
	}
	for pos, char := range str {
		strchar := string(char)
		if strings.Contains(caps, strchar) {
			index := strings.Index(caps, strchar)
			str = str[:pos] + string(caps[(index+number)%26]) + str[pos+1:]
		} else if strings.Contains(lower, strchar) {
			index := strings.Index(lower, strchar)
			str = str[:pos] + string(lower[(index+number)%26]) + str[pos+1:]
		}
	}
	return str
}

func main() {
	fmt.Println(rot(-13, "ATTACK AT DAWN"))
}
