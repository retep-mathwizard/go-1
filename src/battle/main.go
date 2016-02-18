package main

import (
	"fmt"
	i "github.com/retep-mathwizard/utils/src/input"
	c "github.com/skilstak/go/colors"
	"os"
	"reflect"
)

//%s = colors, %f = strings
func getname() (string, string) {
	p1name := i.StringInput("Player 1, what is your name? > ")
	fmt.Println(c.CL)
	p2name := i.StringInput("Player 2, what is your name? > ")
	fmt.Println(c.CL)
	return p1name, p2name
}

//*******************************************************//
type moves struct {
	name1 string
	damage1 int
	accuracy2 int
	name2 string
	damage2 int
	accuracy2 int
}
type class struct {
	moves struct,
	health, agility, shield int
       }

wizard := class{
health: 200,
agility: 50,
shield: 80, 
attacks: moves{
	name1:"fireball"
	damage1:50
	accuracy1:75
	name2:"lightning"
	damage2:125
	accuracy2:25
	}
}
//

ogre := class{
health: 450,
agility: 25,
shield: 10,
attacks: moves{
	name1:"HULK SMASH"
	damage1:500
	accuracy1:10
	name2:"club whack"
	damage2:50
	accuracy2:25
	}
}
//
archer := class{
health: 300,
agility: 100,
shield: 40,
attacks: moves{
	name1:"sniper"
	damage1:25
	accuracy1:100
	name2:"supa-arrow"
	damage2:50
	accuracy2:50
	}
}
//
soldier := class{
health: 400,
agility: 75,
shield: 60,
attacks: moves{
	name1:"broadswipe"
	damage1:50
	accuracy1:70
	name2:"X stroke"
	damage2:20
	accuracy2:100
	}
}
/////////////////////////
func getclass(p1name string, p2name string) (class, class) {
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
