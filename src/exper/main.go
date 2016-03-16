package main

import (
	"fmt"
	"github.com/marcmak/calc/calc"
)

func main() {
	x := calc.Solve("(5*(1+1))/3")
	fmt.Println(x)
}
