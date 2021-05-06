package helpers

import (
	"github.com/sirupsen/logrus"
	"os"
)

func WriteToSet(dependencies map[string]bool, nodes map[string]bool, k string, slice []string) (map[string]bool, map[string]bool) {
	for _, s := range slice {
		if k != s {
			dependencies[k+" --> "+s] = true
		}
		nodes["node "+k] = true
		nodes["node "+s] = true
	}
	return dependencies, nodes
}

func WriteToOsFile(data *os.File, slice map[string]bool) {
	for s := range slice {
		_, err := data.WriteString(s + "\n")

		if err != nil {
			logrus.Fatal(err)
		}
	}
}
