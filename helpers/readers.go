package helpers

import (
	"bufio"
	"github.com/heyjoakim/AARP/logic"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

func Lines(path string) (cr string, slice []string) {
	file, err := os.Open(path)
	if err != nil {
		logrus.Errorf("Opening file causes error: %q\n", err)
	}
	fs := bufio.NewScanner(file)
	dir, f := filepath.Split(path)
	f = f[:len(f)-3]
	tttt := strings.Split(dir, "/")[5]
	for fs.Scan() {
		t, err := logic.Imports(fs.Text())
		if err == nil {
			if strings.HasPrefix(t, "zeeguu") {
				// fmt.Println(t)2or deeper level
				slice = append(slice, strings.Split(t, ".")[0])
			}
		}
	}
	return tttt, slice

}
