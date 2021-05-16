package data

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func WriteFileWithCardinality(data *os.File, slice map[string]int) {
	for s, i := range slice {
		_, err := data.WriteString(s + " : [" + fmt.Sprint(i) + "]" + "\n")

		if err != nil {
			logrus.Fatal(err)
		}
	}
}

func WriteFile(data *os.File, slice map[string]int) {
	for s := range slice {
		_, err := data.WriteString(s + "\n")

		if err != nil {
			logrus.Fatal(err)
		}
	}
}
