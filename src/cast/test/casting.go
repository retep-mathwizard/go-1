package main

import (
	"cast"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	test1 := cast.ParseValue("235")
	test2 := cast.ParseValue("true")
	test3 := cast.ParseValue("1.01")
	tests := []interface{}{test1, test2, test3}
	for num, element := range tests {
		number := strconv.Itoa(num + 1)
		fmt.Println("--------" + number + "---------")
		fmt.Println(element)
		fmt.Println(reflect.TypeOf(test1))
	}
}
