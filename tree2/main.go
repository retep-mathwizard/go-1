package main

import (
	"os"
	"os/exec"
)

func main() {
	where := os.Args[1]
	exec.Command("cd ", where)
	exec.Command("tree > filex.txt")
}
