package main

import (
	"fmt"
	h "github.com/skilstak/go/choice"
	c "github.com/skilstak/go/colors"
	i "github.com/skilstak/go/input"
)

func main() {
	welcome := `
	WELCOME TO THE MAGIC
              _.a$$$$$a._
            ,$$$$$$$$$$$$$.
          ,$$$$$$$$$$$$$$$$$.
         d$$$$$$$$$$$$$$$$$$$b
        d$$$$$$$$~'"'~$$$$$$$$b
       ($$$$$$$p   _   q$$$$$$$)
       $$$$$$$$   (_)   $$$$$$$$
       $$$$$$$$   (_)   $$$$$$$$
       ($$$$$$$b       d$$$$$$$)
        q$$$$$$$$a._.a$$$$$$$$p
         q$$$$$$$$$$$$$$$$$$$p
          '$$$$$$$$$$$$$$$$$'
            '$$$$$$$$$$$$$'
	      '~$$$$$$$~'      BALL`
	fmt.Println(welcome)
	var myArray = []string{
		"Yes",
		"No",
		"Maybe",
		"Certainly",
		"No way!",
		"I don't know...",
	}
	for {
		i.Prompt(c.Rc() + "Ask your question please > " + c.X)
		item := h.Choice(myArray)
		fmt.Println(c.Rc() + c.CL + item + c.X)
	}
}
