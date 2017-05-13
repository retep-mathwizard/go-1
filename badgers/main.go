package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"time"
)

func main() {
	const snake = (`
		 /\ /\         /\ /\         /\ /\              _______   
		/ / \ \       / / \ \       / / \ \             \   _  \  
		/ /   \ \     / /   \ \     / /   \ \            /  /_\  \ 
		/ /    \ \    / /    \ \    / /    \ \           
	/ /     \ \   / /     \ \   / /     \ \           \  \_/   \ 
____________ / /       \ \ / /       \ \ / /       \ \___________\_____  /
/_____/_____/ \/         \/ \/         \/ \/         \/_____/_____/     \/ 
`)
	//const is a constant. Constant's never change.
	const sleepTime = 387
	for i := 1; i <= 4; i++ {
		for i := 1; i <= 12; i++ {
			time.Sleep(sleepTime * time.Millisecond)
			fmt.Println(c.B3 + "badger" + c.X)
		}
		for i := 1; i <= 2; i++ {
			time.Sleep((sleepTime * 2) * time.Millisecond)
			fmt.Println(c.R + "mushroom" + c.X)
		}
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println(snake)
	time.Sleep(500 * time.Millisecond)
	fmt.Println(c.G + "Argh! A Snake!")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println(c.G + "A Snaaaaaaakke!")
	time.Sleep(2500 * time.Millisecond)
	fmt.Println(c.G + "Ooooohhh, its a snake!" + c.X)
	time.Sleep(1500 * time.Millisecond)
	fmt.Println(c.CL)

}
