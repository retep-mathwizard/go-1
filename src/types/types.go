package types

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Float(x interface{}) float64 {
	convertedFloat := 0.1234
	switch x.(type) {
	case string:
		z := x.(string)
		convertedFloat, _ = strconv.ParseFloat(z, 64)
	case int:
		y := x.(int)
		convertedFloat = float64(y)
	case float64:
		convertedFloat = x.(float64)
	default:
		fmt.Println("Invalid type. Terminating program.")
		os.Exit(-1)
	}
	return convertedFloat
}

func Str(x interface{}) string {
	converted := "Not Converted"
	switch x.(type) {
	case int:
		y := x.(int)
		converted = strconv.Itoa(y)
	case string:
		converted = x.(string)
		fmt.Print("")
	case float64:
		z := x.(float64)
		//a type assertion where x is type float64
		converted = strconv.FormatFloat(z, 'f', -1, 64)
		//z is the thing to be turned into a string
		//'f' means decimal *note sigle quotes
		//10 (digits) is decimal precision -1 is all
		//64 refers to float64
	default:
		fmt.Println("Invalid type. Terminating program.")
		os.Exit(-1)
	}
	return converted
}
