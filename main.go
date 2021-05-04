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

func readLines(path string) []string {
  var slice []string
  file, err := os.Open(path)
  if err != nil {
   log.Errorf("Opening file causes error: %q\n", err)
  }
  fs := bufio.NewScanner(file)
  for fs.Scan() {
    t := extractImports(fs.Text())
    for _, s := range t {
     fmt.Println(s)
     slice = append(slice, s)
    }
  }
  return slice

}

func extractImports(line string) []string {
  var slice []string
	x, err := regexp.Compile(`^from ([a-zA-Z]+) import ([a-zA-Z]+)$`)
  if err != nil {
   log.Warnf("Error when compiling regex: %v\n", err)
  }
  res := x.FindAllStringSubmatch(line,-1)
  for i:= 0; i < len(res); i++ {
   for j := 1; j <= 2; j++ {
    fmt.Println(res[i][j])
    slice = append(slice, res[i][j])
   }
  }
  return slice
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
  //fmt.Printf("%q\n ", extractImports("from test import lol")[0][1])
  //fmt.Println(extractImports("from test import looooooool"))
  fmt.Println(readLines("/home/hey/git/Zeeguu-API/zeeguu_core/model/feed.py"))
}
