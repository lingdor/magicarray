package array

import (
	"fmt"
	"github.com/lingdor/magicarray/internal"
	"testing"
)

type benchmarkInfoType struct {
	Field1 string
	Field2 int
}

func GenZvalArray(size int) MagicArray {

	var myarr = internal.EmptyZValArray(true, false, size)
	for i := 0; i < size; i++ {
		myarr = Set(myarr, fmt.Sprintf("i%d", i), i)
	}
	return myarr
}
func GenMap(size int) map[string]any {

	var myarr = make(map[string]any, size)
	for i := 0; i < size; i++ {
		myarr[fmt.Sprintf("i%d", i)] = i
	}
	return myarr
}
func BenchmarkSetMap(b *testing.B) {
	GenMap(b.N)
}
func BenchmarkZval(b *testing.B) {
	GenZvalArray(b.N)
}
