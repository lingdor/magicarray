package utils

import (
	"fmt"
	"strings"
)

func PadLeftRightSpaces(str string, left, full int) string {

	leftSpace := strings.Repeat(" ", left)
	rightSpace := ""
	if len(str)+left < full {
		rightSpace = strings.Repeat(" ", full-len(str)-left)
	}

	return fmt.Sprintf("%s%s%s", leftSpace, str, rightSpace)
}
