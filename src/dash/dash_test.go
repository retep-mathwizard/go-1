package dash_test

import (
	"dash"
	"fmt"
)

func TestDash() {
	test1 := dash.Dasherize(" This is a sentence. ")
	test2 := dash.Dasherize("dash will strip whitespace AND lowercase WoRds")
	test3 := dash.Dasherize("It will then replace all spaces with dashes")
	test4 := dash.DasherizeFile("input.txt")
	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
	dash.WriteLines("output.txt", test4)
}
