package internal

import (
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
)

type TArrayIterator[T any] struct {
	index int
	arr   TArray[T]
}

func (t *TArrayIterator[T]) currentKV() (api.ZVal, api.ZVal) {
	if t.index < t.arr.Len() {
		return zval.NewZValOfKind(kind.Int, t.index), zval.NewZVal(t.arr[t.index])
	}
	return nil, nil
}
func (t *TArrayIterator[T]) NextKV() (api.ZVal, api.ZVal) {
	t.index++
	return t.currentKV()
}

func (t *TArrayIterator[T]) FirstKV() (api.ZVal, api.ZVal) {
	t.index = 0
	return t.currentKV()
}

func (t *TArrayIterator[T]) currentVal() api.ZVal {
	if t.index < t.arr.Len() {
		return zval.NewZVal(t.arr[t.index])
	}
	return nil
}

func (t *TArrayIterator[T]) NextVal() api.ZVal {
	t.index++
	return t.currentVal()
}

func (t *TArrayIterator[T]) FirstVal() api.ZVal {
	t.index = 0
	return t.currentVal()
}
