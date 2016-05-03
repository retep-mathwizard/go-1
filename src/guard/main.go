package main

import (
	"fmt"
	"strings"
)

const one = `
 ██╗
███║
╚██║
 ██║
 ██║
 ╚═╝
`
const two = `
██████╗ 
╚════██╗
 █████╔╝
██╔═══╝ 
███████╗
╚══════╝
`
const three = `
██████╗ 
╚════██╗
 █████╔╝
 ╚═══██╗
██████╔╝
╚═════╝ 
`
const four = `
██╗  ██╗
██║  ██║
███████║
╚════██║
     ██║
     ╚═╝
`
const five = `
███████╗
██╔════╝
███████╗
╚════██║
███████║
╚══════╝
`
const six = `
 ██████╗ 
██╔════╝ 
███████╗ 
██╔═══██╗
╚██████╔╝
 ╚═════╝ 
`
const seven = `
███████╗
╚════██║
    ██╔╝
   ██╔╝ 
   ██║  
   ╚═╝ 
`
const eight = `
 █████╗ 
██╔══██╗
╚█████╔╝
██╔══██╗
╚█████╔╝
 ╚════╝ 
`
const nine = `
 █████╗ 
██╔══██╗
╚██████║
 ╚═══██║
 █████╔╝
 ╚════╝ 
`
const zero = `
 ██████╗ 
██╔═████╗
██║██╔██║
████╔╝██║
╚██████╔╝
 ╚═════╝ 
`

var intToStr = map[string]string{
	"0": zero,
	"1": one,
	"2": two,
	"3": three,
	"4": four,
	"5": five,
	"6": six,
	"7": seven,
	"8": eight,
	"9": nine,
}

func getAscii(number string) [][]string {
	listOfNumbers := strings.Split(number, "")
	var ascii [][]string
	for _, num := range listOfNumbers {
		ascii = append(ascii, strings.Split(intToStr[num], "\n"))
	}
	return ascii
}
func LongestSlice(slices [][]string) int {
	var longest int = 0
	for _, slice := range slices {
		if len(slice) > longest {
			longest = len(slice)
		}
	}
	return longest
}

func joinStrings(stuff [][]string) []string {
	var joined []string
	for lineNumber := 0; lineNumber <= LongestSlice(stuff)-1; lineNumber++ {
		var Oneline string
		for _, item := range stuff {
			Oneline += item[lineNumber]

		}
		joined = append(joined, Oneline)
	}
	return joined
}
func main() {
	stuff := joinStrings(getAscii("2423"))
	for _, line := range stuff {
		fmt.Println(line)
	}
}
