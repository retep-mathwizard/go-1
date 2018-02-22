package main

import (
	"fmt"
	"github.com/skilstak/go-input"
	"math/rand"
	"strings"
	"time"
)

func join_ascii(rletter1, rletter2, rletter3, rcol string) string {
	col := strings.Split(rcol, "\n")
	letter1 := strings.Split(rletter1, "\n")
	letter2 := strings.Split(rletter2, "\n")
	letter3 := strings.Split(rletter3, "\n")
	var result string
	for i := 0; i < len(letter1); i++ {
		result += letter1[i] + col[i] + letter2[i] + col[i] + letter3[i] + "\n"
	}
	return result[:len(result)-1]
}
func drawboard(board [][]int) {
	for i, row := range board {
		if i != 0 {
			fmt.Println(line)
		}
		fmt.Println(join_ascii(letters[row[0]], letters[row[1]], letters[row[2]], col))
	}
}

var letters = map[int]string{2: "\\/\n/\\", 1: "/\\\n\\/", 0: "  \n  "}
var line = "--+--+--"
var col = "|\n|"
var winning_combos = [][]string{{"11", "12", "13"}, {"21", "22", "23"}, {"31", "32", "33"}, {"11", "21", "31"}, {"12", "22", "32"}, {"13", "23", "33"}, {"11", "22", "33"}, {"13", "22", "31"}}
var board = [][]int{{1, 0, 2}, {0, 1, 2}, {0, 0, 1}}

func main() {
	drawboard(board)
	rand.Seed(time.Now().UTC().UnixNano())
	xturn := rand.Intn(2)
	if xturn == 1 {
		input.Ask("where > ")
	}

	fmt.Println(winning_combos)
}
