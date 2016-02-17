package main

import (
	"fmt"
	i "github.com/retep-mathwizard/utils/src/input"
	c "github.com/skilstak/go/colors"
	"os"
	"reflect"
)

//%s = colors, %f = strings
type class struct {
	health, agility, damage, accuracy int
}

func getname() (string, string) {
	p1name := i.StringInput("Player 1, what is your name? > ")
	fmt.Println(c.CL)
	p2name := i.StringInput("Player 2, what is your name? > ")
	fmt.Println(c.CL)
	return p1name, p2name
}
func getclass(p1name string, p2name string) (class, class) {
	wizard := class{health: 200, agility: 50, damage: 80, accuracy: 70}
	ogre := class{health: 450, agility: 25, damage: 100, accuracy: 40}
	archer := class{health: 300, agility: 100, damage: 40, accuracy: 100}
	soldier := class{health: 400, agility: 75, damage: 60, accuracy: 80}
	var sentenceA string = p1name + ",what class? (ogre, wizard, soldier, archer) > "
	var sentenceB string = p2name + " ,what class? (ogre, wizard, soldier, archer) > "
	p1class := soldier
	p2class := soldier
	userclass := i.StringInput(sentenceA)
	switch userclass {
	case "archer":
		p1class = archer
	case "ogre":
		p1class = ogre
	case "wizard":
		p1class = wizard
	case "soldier":
		p1class = soldier
	default:
		fmt.Println("invalid case.")
		os.Exit(-1)
	}
	fmt.Println(c.CL)
	userclass2 := i.StringInput(sentenceB)
	switch userclass2 {
	case "archer":
		p2class = archer
	case "ogre":
		p2class = ogre
	case "wizard":
		p2class = wizard
	case "soldier":
		p2class = soldier
	default:
		fmt.Println("invalid case.")
		os.Exit(-1)
	}
	fmt.Println(c.CL)
	return p1class, p2class
}

func main() {
	p1name, p2name := getname()
	//name
	p1class, p2class := getclass(p1name, p2name)
	//class selection

}
