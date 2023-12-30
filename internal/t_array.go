package internal

import (
	"encoding/json"
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
)

type TArray[T any] []T

func (t TArray[T]) Json() ([]byte, error) {
	return json.Marshal(t)
}

func (t TArray[T]) IsKeys() bool {
	return false
}

func (t TArray[T]) Keys() api.IMagicArray {
	keys := GenListKeys(t.Len())
	return TArray[int](keys)
}

func (t TArray[T]) Values() api.IMagicArray {
	return t
}

func (t TArray[T]) Len() int {
	return len(t)
}

func (t TArray[T]) Get(key interface{}) api.IZVal {
	if index, ok := key.(int); ok {
		if index < t.Len() {
			//todo can improve performance to better in the go newer version
			//use:  zval.NewZValOfKind(t.getKind(), t.arr[index])
			return zval.NewZVal(t[index])
		} else {
			return zval.NewZValNil()
		}
	}
	var zv api.IZVal
	var ok bool

	if zv, ok = key.(api.IZVal); !ok {
		zv = zval.NewZValOfKind(kind.Int, zv)
	}
	if ii, ok := zv.Int(); ok {
		if ii < t.Len() {
			return zval.NewZVal(t[ii])
		} else {
			return zval.NewZValNil()
		}
	}
	return zval.NewZValInvalid()
}

func (t TArray[T]) Iter() api.Iterator {
	return &TArrayIterator[T]{
		arr:   t,
		index: 0,
	}
}
func (t TArray[T]) MarshalJSON() ([]byte, error) {
	return JsonMarshal(t)
}
