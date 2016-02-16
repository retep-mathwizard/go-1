package main

import (
	"fmt"
	"github.com/retep-mathwizard/utils/
)

type class struct {
	health, agility, damage, accuracy int
}

func main() {
	mage := class{health: 200, agility: 50, damage: 80, accuracy: 70}
	ogre := class{health: 450,agility:25,damage:100,accuracy:40}
	dodger := class{health: 300,agility:100,damage:40,accuracy:100}
	soldier := class{health: 400,agility:75,damage:60,accuracy:80}
	
	mage.health = 10
	fmt.Println(mage.health)
}
