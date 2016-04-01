package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	//"github.com/skilstak/go/input"
	"math"
	"math/rand"
	"peter/custom"
	"time"
)

// This is an optimization type thing of https://github.com/retep-mathwizard/fun/blob/master/go/src/heroes/main.go
// All credit goes to @retep-mathwizard on github

func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

type Entity struct {
	Name    string
	Attack  int
	Defense int
	Health  int
	Agility int
}

func CreateAi(entity Entity, name string) Entity {
	aihp := Random(entity.Health, (entity.Health * 2))
	aiagi := Random(1, 10)
	aiatt := int(math.Ceil(float64(aihp) / (float64(aiagi) * 1.5)))
	Ai := Entity{Name: name, Attack: aiatt, Health: aihp, Agility: aiagi}

	return Ai
}

func IsPlayerFirst(agi int, aiagi int) bool {
	if agi >= aiagi {
		return true
	}

	return false
}

func Battle(Attacker Entity, Defender Entity) Entity {
	Defender.Health -= Attacker.Attack
	if Defender.Health <= 0 {
		fmt.Println(c.Clear+c.Y+Attacker.Name+c.B1+" has dealt "+c.R, Attacker.Attack, c.B1+" damage to "+c.Y+Defender.Name+c.B1+" who is now dead")
	} else {
		fmt.Println(c.Clear+c.Y+Attacker.Name+c.B1+" has dealt "+c.R, Attacker.Attack, c.B1+" damage to "+c.Y+Defender.Name+c.B1+" who now has "+c.R, Defender.Health, c.B1+" health")
	}

	return Defender
}

func WhoWon(ai Entity, player Entity) {
	if ai.Health > player.Health {
		fmt.Println(c.Clear + c.Y + ai.Name + c.B1 + " has won!")
	} else {
		fmt.Println(c.Clear + c.Y + player.Name + c.B1 + " has won!")
	}
}

func IsDead(health int) bool {
	if health > 0 {
		return false
	}

	return true
}

var Base Entity = Entity{Name: "Steve", Attack: 6, Health: 20, Agility: 5}
var Player Entity = CreateAi(Base, "Alex")
var Ai Entity = CreateAi(Base, "Scarletwound")
var PlayerFirst bool = IsPlayerFirst(Player.Agility, Ai.Agility)

func main() {
	chars := custom.LoadChars()
	fmt.Println(chars)
	/*
		if PlayerFirst {
			for {
				Ai = Battle(Player, Ai)
				input.Ask(c.B1 + "(Enter To Continue)")
				if IsDead(Ai.Health) {
					break
				}
				Player = Battle(Ai, Player)
				input.Ask(c.B1 + "(Enter To Continue)")
				if IsDead(Player.Health) {
					break
				}
			}
		} else {
			for {
				Player = Battle(Ai, Player)
				input.Ask(c.B1 + "(Enter To Continue)")
				if IsDead(Player.Health) {
					break
				}
				Ai = Battle(Player, Ai)
				input.Ask(c.B1 + "(Enter To Continue)")
				if IsDead(Ai.Health) {
					break
				}
			}
		}
		WhoWon(Ai, Player)
	*/
}
