package main

import (
	"math/rand"
	u "quiz/unicorns"
	"reflect"
	"time"
)

func RandItem(list []func() bool) func() bool {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(len(list))
	randElement := list[randNum]
	return randElement
}

//

func Remove(slice []func() bool, item func() bool) []func() bool {
	var index int

	i2 := reflect.ValueOf(item)
	for p, v := range slice {
		v2 := reflect.ValueOf(v)
		if v2 == i2 {
			index = p
		}
	}
	var result []func() bool
	result = append(result, slice[0:index]...)
	// Append part after the removed element.
	result = append(result, slice[index+1:]...)
	return result
}
func main() {
	questions := []func() bool{u.Q1, u.Q2, u.Q3}
	for questions != nil {
		question := RandItem(questions)
		if question() {
			questions = Remove(questions, question)
		}
	}

}
