package main

import (
	"fmt"
	"strconv"
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

var numbers = []string{zero, one, two, three, four, five, six, seven, eight, nine}

func getAscii(number int) [][]string {
	listOfNumbers := strings.Split(strconv.Itoa(number), "")
	var ascii [][]string
	for _, num := range listOfNumbers {
		intnum, _ := strconv.Atoi(num)
		ascii = append(ascii, strings.Split(numbers[intnum], "\n"))
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

func joinStrings(stuff [][]string) [][]string {
	var joined [][]string
	for lineNumber := 0; lineNumber <= LongestSlice(joined); lineNumber++ {
		var Oneline []string
		for _, item := range stuff {
			//err := item[lineNumber]
			//if err != nil {
			//	panic(err)
			//}
			Oneline = append(Oneline, item[lineNumber])

		}
		joined = append(joined, Oneline)
	}
	return joined
}
func main() {
	stuff := joinStrings(getAscii(2423))
	for _, line := range stuff {
		fmt.Println(line)
	}
}
