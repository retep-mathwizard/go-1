package main

import "fmt"

func remove(slice []string, index int) []string {
	result := []string{}
	result = append(result, slice[0:index]...)
	// Append part after the removed element.
	result = append(result, slice[index+1:]...)
	return result
}

func insert(s []string, at int, val string) []string {
	// Make sure there is enough room
	s = append(s, "0")
	// Move all elements of s up one slot
	copy(s[at+1:], s[at:])
	// Insert the new element at the now free position
	s[at] = val
	return s
}
func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func main() {
	x := []string{"a", "b", "c"}
	fmt.Println(x)
	y := append(x, "d")
	//append to end
	fmt.Println(y)
	//append to beginning
	//first argument to append MUST be a slice second must be a string
	data := append([]string{"BEFORE"}, x...)
	//... lets x (slice) be a string argument
	fmt.Println(data)
	num := []int{1, 2, 3, 4}
	g := append(num, 5, 6)
	fmt.Println(g)
	length := len(g)
	fmt.Println(length)
	//append more than once
	partial := num[2:3]
	fmt.Println(partial)
	list := []string{"a", "b", "c", "d", "d", "e", "f"}
	fmt.Println(list)
	newlist := remove(list, 4)
	fmt.Println(newlist)
	inlist := insert(list, 4, "DDD")
	fmt.Println(inlist)
	dups := []string{"2", "6", "2", "7", "6", "5", "5"}
	fmt.Println(dups)
	nodups := removeDuplicates(dups)
	fmt.Println(nodups)
}
