package main

import (
	"fmt"
	"github.com/reteps/comm/comm"
	"github.com/skilstak/go-input"
	"os"
	"strconv"
	"strings"
	"time"
)

var letters = map[int]string{2: `\\   /
 \\ / 
  X   
 / \\ 
/   \\`, 1: `┌────┐
│    │
│    │
│    │
└────┘`, 0: `      
      
      
      
      `}
var line = "########################\n########################"
var col = `###
###
###
###
###`
var board = [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

func printboard() {
	for i, row := range board {
		if i != 0 {
			fmt.Println(line)
		}
		scol := strings.Split(col, "\n")
		letter1 := strings.Split(letters[row[0]], "\n")
		letter2 := strings.Split(letters[row[1]], "\n")
		letter3 := strings.Split(letters[row[2]], "\n")
		for i := 0; i < len(letter1); i++ {
			fmt.Println(letter1[i] + scol[i] + letter2[i] + scol[i] + letter3[i])
		}
	}
}
func checkwin(board [][]int, code int) bool {
	if code == 0 {
		for _, row := range board {
			for _, char := range row {
				if char == 0 {
					return false
				}
			}
		}
		return true
	}
	var winning_combos = [][]int{{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {2, 4, 6}, {0, 4, 8}}
	for _, combo := range winning_combos {
		if board[(combo[0]-(combo[0]%3))/3][combo[0]%3] == code && board[(combo[1]-(combo[1]%3))/3][combo[1]%3] == code && board[(combo[2]-(combo[2]%3))/3][combo[2]%3] == code {
			return true
		}
	}
	return false
}
func main() {
	connection := comm.Connect(os.Args[1], "1234")
	who := connection.Read()
	xturn := who == "x"
	xwon := false
	owon := false
	tie := false
	var position int
	var err error
	fmt.Println(who, "goes first (you're o)")
	time.Sleep(5 * time.Second)
	for xwon != true && owon != true && tie != true {
		fmt.Println("\033[H\033[2J")
		fmt.Println("You are o, your opponent is x. Pick a number from 1-9 starting at the top left and ending at the bottom right.")
		printboard()
		if xturn {
			fmt.Println("waiting for opponent...")
			position, _ = strconv.Atoi(connection.Read())
			board[((position-1)-(position-1)%3)/3][(position-1)%3] = 2
			if checkwin(board, 2) {
				xwon = true
				fmt.Println("x won!")
			}
			xturn = false
		} else {
			for {
				if checkwin(board, 0) {
					fmt.Println("tie")
					tie = true
					break
				}
				position, err = strconv.Atoi(input.Ask("where (1-9) > "))
				if err == nil && position <= 9 && position >= 1 {
					if board[((position-1)-(position-1)%3)/3][(position-1)%3] == 0 {
						break
					}
				}
				fmt.Println("invalid location")
			}
			connection.Send(strconv.Itoa(position))
			board[((position-1)-(position-1)%3)/3][(position-1)%3] = 1
			if checkwin(board, 1) {
				owon = true
				printboard()
				fmt.Println("o won!")
			}
			xturn = true
		}

	}
	fmt.Println("\033[H\033[2J")
	printboard()

}
