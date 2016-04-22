package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"io"
)
func Dasherize(line string) string {
		return strings.Replace(strings.ToLower(strings.TrimSpace((line))), " ", "-", -1)
}
func DasherizeFile(file string, newfile string) {
	f, _ := os.Open(file)
	// Create a new Scanner for the file.
	var lines []string
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		dashline := Dasherize(line)
		lines = append(lines, dashline)
		
	}
	fmt.Println(lines)
	createdfile, err := os.Create(newfile)
	if err != nil {
		panic(err)
	}
	for _, data := range lines {
		_, err := io.WriteString(createdfile,data + "\n")
		if err != nil {
			panic(err)
		}
	}


}
func main() {
	DasherizeFile("file", "newfile")
}
