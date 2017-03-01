// Echo2 prints its command-line arguments.
package echo2

import (
	"fmt"
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
