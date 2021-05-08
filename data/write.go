package data

import (
	"github.com/sirupsen/logrus"
	"os"
)

func WriteFile(data *os.File, slice map[string]bool) {
	for s := range slice {
		_, err := data.WriteString(s + "\n")

		if err != nil {
			logrus.Fatal(err)
		}
	}
}
