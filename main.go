package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]
	dir := args[0]
	subfiles := getSubFiles(dir)

	var command string = "go build"

	var i int = 0
	for i < len(subfiles) {

		var run bool = true

		if len(subfiles[i]) < 1 {
			run = false
		}

		if run {
			command = command + " " + subfiles[i]
		}

		i++
	}

	fmt.Println(command)
}
