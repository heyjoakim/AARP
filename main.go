package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	log "github.com/sirupsen/logrus"
)

func walker(root string, extension string, exclusion string) *[]string {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if filepath.Ext(path) != extension {
			return nil
		}

		if _, t := filepath.Split(path); t == exclusion {
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		log.Fatalf("Error when walking directories: %v\n", err)
	}

	return &files
}

func countFiles(files *[]string) {
	fmt.Printf("Total number of .py files: %v\n", len(*files))
}

func countLines(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Errorf("Opening file causes error: %q", err)
		return 0
	}

	fs := bufio.NewScanner(file)
	lc := 0
	for fs.Scan() {
		lc++
	}
	return lc
}

func extractImports(line string) {
	x, _ := regexp.Compile(`^from ([a-zA-Z]+) import ([a-zA-Z]+)$`)
	fmt.Println(x)
	res := x.FindAllSubmatch([]byte(line), -1)[0][1]
	fmt.Println(string(res))
}

func main() {
	root := "/home/hey/git/Zeeguu-API"
	extension := ".py"
	exclusion := "__init__.py"
	test := walker(root, extension, exclusion)
	countFiles(test)

	for _, j := range *test {
		fmt.Println(countLines(j), j)
	}
	extractImports("from lol import test")
}
