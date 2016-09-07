package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	"github.com/skilstak/go-input"
	"strconv"
)

func main() {
	number, err := strconv.Atoi(input.Ask(c.B3 + "Pick a number > "))
	if err != nil {
		fmt.Println("Thats not a number...")
		panic(err)
	}
	for multiplier := 1; multiplier <= 20; multiplier++ {
		answer := number * multiplier
		fmt.Printf("%v * %v = %v\n", multiplier, number, answer)
		//%s for colors, %v for variables
	}
	fmt.Print(c.X)
}
