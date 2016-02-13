package works

import (
	"fmt"
	"strings"
)

//for loop
//   if strings.contains(element 1)
//   else
//
//
func arrayContains(stuffToCheckFor []string, stuff string) bool {
	mode := false
	for _, substr := range stuffToCheckFor {
		if strings.Contains(stuff, substr) {
			mode = true
		}
	}
	if mode == true {
		return true
	} else {
		return false
	}
}
func arrayEquals(stuffToCheckFor []string, stuff string) bool {
	mode := false
	for _, str := range stuffToCheckFor {
		if stuff == str {
			mode = true
		}
	}
	if mode == true {
		return true
	} else {
		return false
	}
}
func main() {
	//slice
	slicenames := []string{"bob", "jim", "manny", "sally", "bo"}
	//array has a size defined
	arraynames := [3]string{
		"leto",
		"paul",
		"teg",
	}
	fmt.Println(arraynames)
	for _, element := range slicenames {
		//_ is the index, slicenames
		//_ indaicates not used
		//element is the item in the index
		//range is the length of an slice/array
		fmt.Println(element)
	}
	hername := "sammy"
	inside := arrayContains(slicenames, hername)
	fmt.Println(inside)
	equals := arrayEquals([]string{"joe", "bo"}, "bo")
	fmt.Println(equals)
}
