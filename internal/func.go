package internal

import (
	"github.com/lingdor/magicarray/api"
	"reflect"
)

var JsonMarshal func(arr api.IMagicArray, opts ...api.JsonOpt) ([]byte, error)

func GenListKeys(len int) []int {
	var keys = make([]int, len)
	for i := 0; i < len; i++ {
		keys[i] = i
	}
	return keys
}

func SlicetoAnyList(refVal reflect.Value) []any {
	len := refVal.Len()
	ret := make([]any, 0, len)
	for i := 0; i < len; i++ {
		ret = append(ret, refVal.Index(i).Interface())
	}
	return ret
}
func newTArray[T any](listVal []T) api.IMagicArray {
	return TArray[T](listVal)
}
