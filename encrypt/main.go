package main

import (
	"fmt"
	"strconv"
	"strings"
)

func moveToEnd(array []string) []string {
	return append(array[1:], array[0])
}
func findPos(array []string, value string) string {
	for pos, element := range array {
		if element == value {
			return strconv.Itoa(pos)
		}
	}
	return "well, that failed"
}
func main() {
	const trash = "ZWQJDAPDIV"
	num := "328827109"
	fmt.Println(num)
	var encrypted string
	array := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	arrayCopy := array
	for _, char := range num {
		encrypted += findPos(array, string(char))
		array = moveToEnd(array)
	}
	var newtrash string
	for _, n := range encrypted {
		i, _ := strconv.Atoi(string(n))
		newtrash += string(trash[i])
	}
	fmt.Println(newtrash)
	//-------decode------
func decode(newtrash string) string {
	var decoded string
	for _, letter := range newtrash {
		orig := strconv.Itoa(strings.Index(trash, string(letter)))
		decoded += orig
	}

	var decrypted string
	for _, char := range strings.Split(decoded, "") {
		i, _ := strconv.Atoi(char)
		oldChar := arrayCopy[i]
		arrayCopy = moveToEnd(arrayCopy)
		decrypted += oldChar
	}
	fmt.Println(decrypted)
}
}
