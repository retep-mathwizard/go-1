package main

import (
	"github.com/retep-mathwizard/fun/go/src/heroes/bmath"
	"github.com/retep-mathwizard/fun/go/src/heroes/engine"
	"github.com/retep-mathwizard/fun/go/src/heroes/info"
	"github.com/retep-mathwizard/fun/go/src/heroes/load"
	o "github.com/retep-mathwizard/utils/other"
	//c "github.com/skilstak/go/colors"
)

func main() {
	p1Name, p2Name := info.GetName()
	//name
	p1Class := load.LoadClass("../data/class.json")
	#ai#p2Class := load.LoadClass("../data/class2.json")
	//class selection
	p2HasShield := false
	p2MovePos := 1
	p1Health := p1Class.Health
	p2Health := p2Class.Health
	for {
		p1MovePos, p1HasShield := engine.ChooseMove(p1Name, p1Class)
		//get move
		o.Spacer(3)
		p2Health = engine.Attack(p1Class, p1Name, p2Class, p2Name, p2Health, p1MovePos, p2HasShield)
		bmath.EvalAlive(p2Name, p1Name, p2Health)
		o.Spacer(3)
		//attack
		//remove the shield, turn is over
		#ai#p2MovePos, p2HasShield = engine.ChooseMove(p2Name, p2Class)
		o.Spacer(3)
		//p1class with the name of p1name attacks p2class with the name of p2name and the health of p2health using move #movePos, damage is reduced if p2 has a shield.
		p1Health = engine.Attack(p2Class, p2Name, p1Class, p1Name, p1Health, p2MovePos, p1HasShield)
		bmath.EvalAlive(p1Name, p2Name, p1Health)
		o.Spacer(3)

	}
}
