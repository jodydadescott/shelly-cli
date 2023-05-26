package util

import (
	"regexp"
)

const (
	hexadecimal string = "^[0-9a-fA-F]+$"
)

var (
	rxHexadecimal = regexp.MustCompile(hexadecimal)
)

// IsHexadecimal check if the string is a hexadecimal number.
func IsHexadecimal(str string) bool {
	return rxHexadecimal.MatchString(str)
}
