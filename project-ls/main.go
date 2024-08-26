package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Your solution goes here. Good luck!
	dirname := "testdata"
	files := listFiles(dirname)

	for _, f := range files {
		fmt.Println(f)
	}

	//fmt.Println(strings.Join(files, " "))
}

func listFiles(dirname string) []string {
	var dirs []string

	files, err := os.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	// dno should be true or false
	var allFiles = flag.Bool("a", false, "show all files")
	flag.Parse()

	for _, f := range files {
		if strings.HasPrefix(f.Name(), ".") && !*allFiles {
			continue
		}

		dirs = append(dirs, f.Name())
	}

	return dirs
}
