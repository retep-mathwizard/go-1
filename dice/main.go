package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	"github.com/skilstak/go-input"
	"github.com/skilstak/go-random"
)

func main() {

	const (
		side1 = (`
 ------
|      |
|  壹  |
|     1|
 ------
		`)

		side2 = (`
 ------
|      |
|  貳  |
|     2|
 ------
		`)

		side3 = (`
 ------
|      |
|  叁  |
|     3|
 ------
		`)

		side4 = (`
 ------
|      |
|  肆  |
|     4|
 ------
 		`)

		side5 = (`
 ------
|      | 
|  伍  |
|     5|
 ------
		`)

		side6 = (`
 ------
|      | 
|  陸  |
|     6|
 ------
		`)
	)
	sides := []string{side1, side2, side3, side4, side5, side6}
	for {
		randomSide := random.Choice(sides)
		fmt.Print(c.CL + c.Rc() + randomSide + c.X)
		input.Ask("")
	}
}
