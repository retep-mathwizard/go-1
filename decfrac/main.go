package main

import (
	"fmt"
	"github.com/retep-mathwizard/utils/convert"
	"github.com/retep-mathwizard/utils/input"
	"math"
	"strings"
)

func main() {
	fmt.Println("Input your decimal")
	num := input.StringInput(" > ")
	digits := num[strings.Index(num, ".")+1:]
	decimal := num[strings.Index(num, "."):]
	numbers := num[:strings.Index(num, ".")]
	denom := convert.FloatToInt(math.Pow(10, convert.IntToFloat(len(digits))))
	numer := convert.FloatToInt(convert.StringToFloat(decimal) * convert.IntToFloat(denom))
	var gcf int
	for i := numer/2 + 1; i >= 0; i-- {
		if (numer%i == 0) && (denom%i == 0) {
			gcf = i
			break
		}

	}
	newnum := numer / gcf
	newdum := denom / gcf
	fmt.Println(numbers+" and ", newnum)
	fmt.Println("    -------")
	fmt.Println("      ", newdum)
}
