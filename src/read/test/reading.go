package main

import (
	"fmt"
	"read"
)

func main() {
	for _, line := range read.GetContents("filetest") {
		fmt.Println(line)
	}
}
