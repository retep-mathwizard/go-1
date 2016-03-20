package main

import (
	"fmt"
	"github.com/marcmak/calc/calc"
)

func main() {
	x := calc.Solve("SQRT(5^3)")
	fmt.Println(x)
}
