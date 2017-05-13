package main

import (
	"fmt"
	"github.com/retep-mathwizard/utils/input"
	"github.com/retep-mathwizard/utils/mod"
	c "github.com/skilstak/go/colors"
	"strings"
)

func StringToSlice(str string) []string {
	stringSlice := strings.Split(str, "|")
	return stringSlice
}
func SliceToString(array []string) string {
	str := strings.Join(array, "")
	return str

}
func GetResponse(EasterEggs [][]string, Answers []string, input string) string {

	for _, item := range EasterEggs {
		//each easter egg list
		responses := item[0]
		slicedResponses := StringToSlice(responses)
		for _, stuff := range slicedResponses {
			//each easter egg
			if strings.Contains(input, stuff) {
				return item[1]
			}
		}
	}
	return mod.RandStringItem(Answers)
}
func main() {
	Death := []string{"die|death|kill|murder", "Why are you thinking suck dark thoughts?"}
	Love := []string{"kiss|love", "Uhhhhhhhhhhh"}
	Why := []string{"why|who", "I was wondering the same question"}
	EasterEggs := [][]string{Death, Love, Why}
	Answers := []string{"Yep", "Nope", "Are you kidding me?", "What do you think?", "It is decidedly so.", "Without a doubt.", "Yes definitely.", "You may rely on it.", "As I see it, yes.", "Most likely.", "Outlook good.", "Yes.", "Signs point to yes.", "Reply hazy try again.", "Ask again later.", "Better not tell you now.", "Cannot predict now.", "Concentrate and ask again.", "Don't count on it.", "My reply is no.", "My sources say no.", "Outlook not so good.", "Very doubtful."}
	for {
		user := strings.ToLower(strings.TrimSpace(input.StringInput(c.Rc() + " > " + c.X)))
		response := GetResponse(EasterEggs, Answers, user)
		fmt.Println(c.B3 + response + c.X)
	}
}
