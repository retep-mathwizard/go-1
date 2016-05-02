package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const trash = "ZWQJDAPDIV"
	var numbers = "23039285723"
	var newtrash string
	for _, num := range numbers {
		i, _ := strconv.Atoi(string(num))
		newtrash += string(trash[i])
	}
	fmt.Println(numbers)
	fmt.Println(newtrash)
	var decoded string
	for _, letter := range newtrash {
		orig := strconv.Itoa(strings.Index(trash, string(letter)))
		decoded += orig
	}
	fmt.Println(decoded)
}
