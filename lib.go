package main

import (
	"fmt"
	"os"
	"strings"
)

func getSubFiles(path string) []string {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	names, err := file.Readdirnames(0)
	if err != nil {
		fmt.Println(err)
	}

	var i int = 0
	for i < len(names) {

		var currentItem = names[i]

		if !strings.Contains(currentItem, ".go") {
			names[i] = ""
		}

		i++
	}

	return names
}
