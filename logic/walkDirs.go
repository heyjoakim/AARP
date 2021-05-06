package logic

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func Walk(root string, extension string, exclusion string) *[]string {
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
		logrus.Fatalf("Error when walking directories: %v\n", err)
	}

	return &files
}
