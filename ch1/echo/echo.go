package echo

import (
	"strings"
	"os"
)

func Echo2() string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func Echo3() string {
	return strings.Join(os.Args[1:], " ")
}
