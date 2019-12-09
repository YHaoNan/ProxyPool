package utils

import "strings"

func TrimSpaceCharacter(text string) string {
	return strings.Replace(strings.Replace(strings.Replace(text, " ", "", -1), "\n", "", -1), "\t", "", -1)
}

