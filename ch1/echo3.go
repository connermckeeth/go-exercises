// Echo3 prints the command line arguments from the os
package echo3

import (
	"fmt"
	"os"
	"strings"
)

func Echo3() string {
	return strings.Join(os.Args[1:], " ")
}
