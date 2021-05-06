package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

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

func readLines(path string) (cr string, slice []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Errorf("Opening file causes error: %q\n", err)
	}
	fs := bufio.NewScanner(file)
	dir, f := filepath.Split(path)
	f = f[:len(f)-3]
	tttt := strings.Split(dir, "/")[5]
	for fs.Scan() {
		t, err := extractImports(fs.Text())
		if err == nil {
			if strings.HasPrefix(t, "zeeguu") {
				// fmt.Println(t)2or deeper level
				slice = append(slice, strings.Split(t, ".")[0])
			}
		}
	}
	return tttt, slice

}

func WriteToSet(sset map[string]bool, lset map[string]bool, k string, slice []string) (map[string]bool, map[string]bool) {

	for _, s := range slice {
		if k != s {
			sset[k+" --> "+s] = true
		}
		lset["node "+k] = true
		lset["node "+s] = true
	}
	return sset, lset

}

func writeToSlice(data *os.File, slice map[string]bool) {

	for s := range slice {
		_, err := data.WriteString(s + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}
}

func extractImports(line string) (stringValue string, err error) {
	ra := regexp.MustCompile(`^import (\S+)`)
	rb := regexp.MustCompile(`^from (\S+)`)

	x := ra.FindAllStringSubmatch(line, -1)
	y := rb.FindAllStringSubmatch(line, -1)

	if y != nil {
		return y[0][1], nil
	} else if x != nil {
		return x[0][1], nil
	}

	return "", errors.New("No imports found")
}

func main() {
	//root := "/home/hey/git/multiparty_computation"
	root := "/home/hey/git/Zeeguu-API"

	extension := ".py"
	exclusion := "__init__.py"
	test := walker(root, extension, exclusion)
	countFiles(test)

	f, err := os.Create("data.wsd")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("@startuml\ntitle Automated Diagram\nskinparam nodesep 100\nskinparam ranksep 100\n")
	set := make(map[string]bool)
	fset := make(map[string]bool)

	for _, j := range *test {
		newf, sl := readLines(j)
		set, fset = WriteToSet(set, fset, newf, sl)
	}
	writeToSlice(f, fset)
	writeToSlice(f, set)
	f.WriteString("@enduml")
	defer f.Close()

}
