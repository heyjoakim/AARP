package logic

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

func ReadLines(path string) (cr string, slice []string) {
	file, err := os.Open(path)
	if err != nil {
		logrus.Errorf("Opening file causes error: %q\n", err)
	}
	fs := bufio.NewScanner(file)
	dir, _ := filepath.Split(path)
	dirLevel := strings.Split(dir, "/")[5]

	for fs.Scan() {
		t, err := SearchImports(fs.Text())
		if err == nil {
			if strings.HasPrefix(t, "zeeguu") {
				slice = append(slice, strings.Split(t, ".")[0])
			}
		}
	}
	return dirLevel, slice
}
