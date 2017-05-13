package bmath

import (
	"fmt"
	o "github.com/retep-mathwizard/utils/other"
	c "github.com/skilstak/go/colors"
	"math"
	"math/rand"
	"time"
)

func Round(f float64) int {
	return int(math.Floor(f + .5))
}
func EvalAlive(name string, opponentsname string, health int) {
	if health <= 0 {
		fmt.Println(c.CL+"Oh no!"+name+"now has", health, ". Looks like", opponentsname, "won!")
		o.Exit("")
	}

}
func Roll(accuracy int) bool {
	rand.Seed(time.Now().UnixNano())
	var randRoll int
	randRoll = rand.Intn(100)
	if randRoll > accuracy {
		return false
	} else {
		return true
	}

}

func DptCalc(player *load.Race) {
	var dptlist []int
	for _, attack := range player.Attacks {
		damage := attack.Damage
		acc := attack.Accuracy
		var dpt float64 = float64(damage) * (float64(acc) / 100)
		dptlist = append(dptlist, Round(dpt))
	}
}
