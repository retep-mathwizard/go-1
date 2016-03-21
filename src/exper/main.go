package main

import (
	"fmt"
	"github.com/retep-mathwizard/utils/convert"
	"math"
	"strings"
)

func main() {
	num := 0.034
	newnum := convert.FloatToString(num)
	digits := newnum[strings.Index(newnum, ".")+1:]
	fmt.Println(math.Pow(10, convert.IntToFloat(len(digits))))
}
