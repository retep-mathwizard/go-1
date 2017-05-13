package main

import (
	"fmt"
	"github.com/retep-mathwizard/utils/mmath"
	c "github.com/skilstak/go/colors"
)

func Space(num int) {
	x := 0
	for x != num {
		fmt.Print(" ")
		x += 1
	}

}
func main() {
	for {
		randSpace := mmath.RandInt(1, 5)
		fmt.Print(c.Rc() + "Nyan" + c.X)
		Space(randSpace)
	}
}
