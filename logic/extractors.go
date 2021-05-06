package logic

import (
	"errors"
	"regexp"
)

func Imports(line string) (stringValue string, err error) {
	ra := regexp.MustCompile(`^import (\S+)`)
	rb := regexp.MustCompile(`^from (\S+)`)

	x := ra.FindAllStringSubmatch(line, -1)
	y := rb.FindAllStringSubmatch(line, -1)

	if y != nil {
		return y[0][1], nil
	} else if x != nil {
		return x[0][1], nil
	}

	return "", errors.New("no imports found")
}

