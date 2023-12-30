package array

import (
	"bytes"
	"strings"
)

// Implode to join the values to a string
func Implode(arr MagicArray, separatorChar string) string {
	iter := arr.Iter()
	buffer := bytes.Buffer{}
	for v := iter.FirstVal(); v != nil; v = iter.NextVal() {
		if iter.Index() > 0 {
			buffer.Write([]byte(separatorChar))
		}
		buffer.Write([]byte(v.String()))
	}
	return string(buffer.Bytes())
}

// Explode string to MagicArray
func Explode(str string, separatorChar string) MagicArray {
	splitStrs := strings.Split(str, separatorChar)
	return ValueOfSlice(splitStrs)
}
