package main

import (
	"fmt"
	"github.com/retep-mathwizard/utils/convert"
	"github.com/retep-mathwizard/utils/input"
	"math"
	"strings"
)

func main() {
	num := input.StringInput(" > ")
	digits := num[strings.Index(num, ".")+1:]
	decimal := num[strings.Index(num, "."):]
	numbers := num[:strings.Index(num, ".")]
	denom := math.Pow(10, convert.IntToFloat(len(digits)))
	numer := convert.StringToFloat(decimal) * denom
	intnumer := convert.FloatToInt(numer * denom)
	intdenom := convert.FloatToInt(denom)
	fmt.Println(denom)
	fmt.Println(numer)
	var factors []int
	for i := 1; i <= convert.FloatToInt(numer); i++ {
		if (intnumer%i == 0) && (intdenom%i == 0) {
			factors = append(factors, i)
		}

	}
	dividend := factors[len(factors)-1]
	newnumer := denom / dividend
	newdenom := numer / dividend
	fmt.Println(numbers+"and  ", newnumer)
	fmt.Println("             -------")
	fmt.Println("           ", newdenom)
}
