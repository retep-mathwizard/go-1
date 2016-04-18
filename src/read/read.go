package read

import (
	"bufio"
	"os"
	"strings"
)

func GetContents(file string) []string {
	f, _ := os.Open(file)
	// Create a new Scanner for the file.
	var lines []string
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		strippedline := strings.TrimSpace(line)
		if strings.HasPrefix(strippedline, "#") != true {
			lines = append(lines, strippedline)
		}
	}
	return lines

}
