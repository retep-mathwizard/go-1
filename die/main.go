package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	numDice = 3
)

func StringToSlice(str string, splitat string) []string {
	stringSlice := strings.Split(str, splitat)
	return stringSlice
}
func RandStringItem(list []string) string {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(len(list))
	randElement := list[randNum]
	return randElement
}
func main() {
	const (
		side1 = (` ------ 
|      |
|  壹  |
|     1|
 ------
		`)

		side2 = (` ------ 
|      |
|  貳  |
|     2|
 ------ 
		`)

		side3 = (` ------ 
|      |
|  叁  |
|     3|
 ------ 
		`)

		side4 = (` ------ 
|      |
|  肆  |
|     4|
 ------ 
 		`)

		side5 = (` ------ 
|      | 
|  伍  |
|     5|
 ------ 
		`)

		side6 = (` ------ 
|      | 
|  陸  |
|     6|
 ------ 
		`)
	)
	dicelist := []string{side1, side2, side3, side4, side5, side6}
	var dice [][]string
	for i := 0; i < numDice; i++ {
		die := StringToSlice(RandStringItem(dicelist), "\n")
		dice = append(dice, die)
	}
	for i := 0; i < 5; i++ {
		for j, n := 0, numDice; j < n; j++ {

			fmt.Print(dice[j][i])

		}
		fmt.Println()
	}

}
