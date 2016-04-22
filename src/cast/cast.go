package cast

import (
	"reflect"
	"strconv"
	"strings"
)

func ParseValue(unstrippedValue string) interface{} {
	value := strings.TrimSpace(unstrippedValue)
	num, err := strconv.Atoi(value)
	if err == nil {
		return num
	}

	boo, err := strconv.ParseBool(value)
	if err == nil {
		return boo
	}
	dec, err := strconv.ParseFloat(value, 64)
	if err == nil {
		return dec
	}
	if reflect.TypeOf(value) == reflect.TypeOf("a string") {
		return value
	} else {
		panic(err)
	}
}
