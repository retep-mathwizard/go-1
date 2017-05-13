package main

import (
	"fmt"
	"github.com/retep-mathwizard/utils/input"
	c "github.com/skilstak/go/colors"
)

func main() {
	waffles := input.StringInput(c.Y + "Do you like waffles? > ")
	if waffles == "yes" {
		pancakes := input.StringInput(c.Y + "Do you like pancakes? > ")
		if pancakes == "yes" {
			toast := input.StringInput(c.Y + "Do you like french toast? > ")
			if toast == "yes" {
				fmt.Println(c.M + "♫ do do do Can't wait to get a mouthful do do do ♫" + c.X)

			} else if toast == "no" {
				fmt.Println(c.R + "aww" + c.X)
			} else {
				fmt.Println(c.C + "Start speaking some english..." + c.X)
			}
		} else if pancakes == "no" {
			fmt.Println(c.R + "awww" + c.X)
		} else {
			fmt.Println(c.C + "Start speaking some english..." + c.X)
		}
	} else if waffles == "no" {
		fmt.Println(c.R + "awww" + c.X)
	} else {
		fmt.Println(c.C + "Start speaking some english..." + c.X)
	}
}
