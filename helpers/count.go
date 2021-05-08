package helpers

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func CountFiles(files *[]string) {
	fmt.Printf("Total number of .py files: %v\n", len(*files))
}

func CountLines(path string) int {
	file, err := os.Open(path)
	if err != nil {
		logrus.Errorf("Opening file causes error: %q", err)
		return 0
	}

	fs := bufio.NewScanner(file)
	lc := 0
	for fs.Scan() {
		lc++
	}
	return lc
}
