package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func walkFn(files *[]string, extension string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("Error when processing walker function: %v\n", err)
		}
		if filepath.Ext(path) != extension {
			return nil
		}

		*files = append(*files, path)
		return nil
	}
}

func walkOut(root string, extension string) {
	var files []string
	err := filepath.Walk(root, walkFn(&files, extension))
	if err != nil {
		log.Fatalf("Error when walking directories: %v\n", err)
	}

	for _, f := range files {
		fmt.Printf("%v\n", f)
	}
	fmt.Printf("Total number of .py files: %v\n", len(files))
}

func countLines(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Errorf("Opening file causes error: %q", err)
		return 0, nil
	}

	fs := bufio.NewScanner(file)
	lc := 0
	for fs.Scan() {
		lc++
	}
	return lc, nil
}

func main() {
	root := "/home/hey/git/Zeeguu-API%"
	//extension := ".py"
	// walkOut(root, extension)
	countLines(root)
}
