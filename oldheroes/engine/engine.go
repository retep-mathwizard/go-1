package engine

import (
	"fmt"
	i "github.com/retep-mathwizard/utils/input"
	o "github.com/retep-mathwizard/utils/other"
	m "oldheroes/bmath"
	//c "github.com/skilstak/go/colors"
	"oldheroes/load"
	"strings"
)

func Attack(attacker *load.Race, attackername string, defender *load.Race, defendername string, health int, movePos int, hasShield bool) int {
	//fmt.Println(willHit)
	var newHealth int
	if movePos > -1 {
		willHit := m.Roll(attacker.Attacks[movePos].Accuracy)
		//returns true of false, based on accuracy of the move
		//if the attack isn't sheilding:
		if hasShield != true {
			//if the opponent doesn't have a sheild
			if willHit {
				//if it hit's:
				newHealth = health - attacker.Attacks[movePos].Damage
				fmt.Println("You caused", attacker.Attacks[movePos].Damage)
				fmt.Println("The defender now has", newHealth)
			} else {
				fmt.Println("Oh No!" + attackername + " missed his foe!")
				newHealth = health
			}
		} else {
			if willHit {
				var shield float64 = 90
				var percentage float64 = shield * 0.01
				initialDamage := attacker.Attacks[movePos].Damage
				blockedDamage := percentage * float64(initialDamage)
				remainingDamage := float64(initialDamage) - blockedDamage
				newHealth = m.Round(float64(health) - remainingDamage)
				fmt.Println("Success! " + attackername + " attacked " + defendername + " , but" + defendername + " shielded.")
				fmt.Println("The  shield blocked ", blockedDamage, " damage, but you still caused", remainingDamage)
				fmt.Println("The defender now has", newHealth)
			} else {
				fmt.Println("Oh No!" + defendername + " shielded, but " + attackername + " missed his foe!")
				newHealth = health

			}
		}
	} else {
		fmt.Println("You sheilded, so you dealt no damage.")
		newHealth = health
	}
	return newHealth
}
func InputMove() string {
	move := strings.ToLower(i.StringInput("Name > "))
}
func IsValid(class *load.Race, move string) bool {
	for j, _ := range class.Attacks {
		stringName := strings.ToLower(string(class.Attacks[j].Name))
		if move == stringName {
			return true
		}
		return false
	}
}
func ChooseMove(name string, class *load.Race) (int, bool) {
	fmt.Println("Choose your move " + name)
	for i, _ := range class.Attacks { // Prints "Jonny" and "Jenna"
		fmt.Println("Name:", class.Attacks[i].Name)
		fmt.Println("Info:", class.Attacks[i].Info)
		fmt.Println("Damage:", class.Attacks[i].Damage)
		fmt.Println("Accuracy:", class.Attacks[i].Accuracy)
		o.Spacer(1)
	}
	fmt.Println("Name: Shield")
	fmt.Println("Percent of damage blocked: 90%")
	movePos := -1
	validMove := false
	hasShield := false
	for validMove == false {
		move := InputMove()
		if IsValid(class, move) != true {
			if move == "sheild" || move == "shield" {
				validMove = true
				hasShield = true
				fmt.Println("Move selected!")
			} else {
				fmt.Println("Invalid move.")
			}
		} else {

		}
	}
	return movePos, hasShield
}
