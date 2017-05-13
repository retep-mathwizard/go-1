package main

import (
	"fmt"
	"github.com/skilstak/go/input"
	"math"
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

type Player struct {
	Name    string
	Attack  int
	Health  int
	Agility int
}

func createAi(player Player) Player {
	aihp := random(player.Health, (player.Heath * 2))
	aiagi := random(1, 10)
	aiatt := int(math.Ceil(float64(aihp) / (float64(aiagi) * 1.5)))
	Ai := Player{Name: "Evil Boss", Attack: aiatt, Health: aihp, Agility: aiagi}
	return Ai
}
func WhoGoesFirst(playagi int, aiagi int) bool {
	if playagi >= aiagi {
		return true
	}
	return false
}
func battle(Attacker Player, Defender Player) Player {
	Defender.Health -= Attacker.Attack
	if Defender.Health <= 0 {
		fmt.Println(c.Clear+c.Y+Attacker.Name+c.B1+" has dealt "+c.R, Attacker.Attack, c.B1+" damage to "+c.Y+Defender.Name+c.B1+" who is now dead")
	} else {
		fmt.Println(c.Clear+c.Y+Attacker.Name+c.B1+" has dealt "+c.R, Attacker.Attack, c.B1+" damage to "+c.Y+Defender.Name+c.B1+" who now has "+c.R, Defender.Health, c.B1+" health")
	}
	return Defender
}
func WhoWon(ai Player, pl Player) {
	if ai.Health > player.Health {
		fmt.Println(c.Clear + c.Y + ai.Name + c.B1 + " has won!")
	} else {
		fmt.Println(c.Clear + c.Y + player.Name + c.B1 + " has won!")
	}
}
func isDead(x int) bool {
	if x > 0 {
		return false
	}
	return true
}
func main() {
	Character := Player{Name: "Bob", Attack: 12, Health: 100, Agility: 15}
	Ai := createAi(Character)
	alive := false
	PlayerGoes1st := WhoGoesFirst(Character.Agility, Ai.Agility)
	if PlayerGoes1st {
		for {
			Ai = battle(Character, Ai)
			input.Ask(c.B1 + "(Enter To Continue)")
			if isDead(Ai.Health) {
				break
			}
			Character = battle(Ai, Character)
			input.Ask(c.B1 + "(Enter To Continue)")
			if isDead(Character.Health) {
				break
			}
		}
	} else {
		for {
			Character = battle(Ai, Character)
			input.Ask(c.B1 + "(Enter To Continue)")
			if isDead(Character.Health) {
				break
			}
			Ai = battle(Character, Ai)
			input.Ask("")
			if isDead(Ai.Health) {
				break
			}
		}
	}
	WhoWon(Ai, Character)

}
