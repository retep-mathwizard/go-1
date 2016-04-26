package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Dasherize(line string) string {
	return strings.Replace(strings.ToLower(strings.TrimSpace((line))), " ", "-", -1)
}
func DasherizeFile(file string) []string {
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
	return lines

}
func WriteLines(file string, data []string) {
	createdfile, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	for _, line := range data {
		_, err := io.WriteString(createdfile, line+"\n")
		if err != nil {
			panic(err)
		}
	}
}
func main() {
	DasherizeFile("file", "newfile")
}
