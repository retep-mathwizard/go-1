package main

import (
	"fmt"
	"github.com/marcmak/calc/calc"
	"github.com/retep-mathwizard/utils/input"
	"strings"
)

type AllFormulas struct {
	formulas []Formula
}
type Formula struct {
	name     string
	equation string
	vars     []string
}

func geteq(stuff AllFormulas) int {
	eq := input.StringInput("What is the name of the formula? > ")
	neweq := strings.TrimSpace(strings.ToLower(eq))
	for _, elem := range stuff.formulas {
		if elem.name == neweq {
			return findpos(stuff.formulas, elem.name)
		}
	}
	return -1
}
func findpos(formulas []Formula, elem string) int {
	for p, v := range formulas {
		if v.name == elem {
			return p
		}
	}
	return -1
}
func varreplace(formula string) (string, string) {
	i := strings.Index(formula, "=")
	rsof := formula[i+1:]
	lsof := formula[:i]
	vars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	fmt.Println(formula)
	for _, item := range vars {
		if strings.Contains(rsof, item) {
			value := input.StringInput(fmt.Sprintf("Input the value of %s. > ", item))
			rsof = strings.Replace(rsof, item, value, -1)
		}
	}
	return rsof, lsof
}
func printvars(pos int, stuff AllFormulas) {
	list := stuff.formulas[pos].vars
	for _, item := range list {
		fmt.Println(item)
	}
}

func main() {
	//i := strings.Index(x, "@")
	test := Formula{
		name:     "testing",
		equation: "y=x+(5*z)-1",
		vars:     []string{"x is the x value", "y = total"},
	}
	py := Formula{
		name:     "py",
		equation: "c=SQRT(a*a + b*b)",
		vars:     []string{},
	}

	stuff := AllFormulas{formulas: []Formula{test, py}}
	pos := geteq(stuff)
	var formula string
	if pos < 0 {
		formula = input.StringInput("Input the formula > ")
	} else {
		printvars(pos, stuff)
		formula = stuff.formulas[pos].equation
	}
	newformula, lsoe := varreplace(formula)
	answer := calc.Solve(newformula)
	fmt.Println("Solving " + formula + ". ")
	fmt.Println(lsoe+" =", answer)

}
