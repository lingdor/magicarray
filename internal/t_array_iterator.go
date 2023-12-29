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

func (t *TArrayIterator[T]) Index() int {
	return t.index
}

func (t *TArrayIterator[T]) currentKV() (api.IZVal, api.IZVal) {
	if t.index < t.arr.Len() {
		return zval.NewZValOfKind(kind.Int, t.index), zval.NewZVal(t.arr[t.index])
	}
	return nil, nil
}
func (t *TArrayIterator[T]) NextKV() (api.IZVal, api.IZVal) {
	t.index++
	return t.currentKV()
}

func (t *TArrayIterator[T]) FirstKV() (api.IZVal, api.IZVal) {
	t.index = 0
	return t.currentKV()
}

func (t *TArrayIterator[T]) currentVal() api.IZVal {
	if t.index < t.arr.Len() {
		return zval.NewZVal(t.arr[t.index])
	}
	return nil
}

func (t *TArrayIterator[T]) NextVal() api.IZVal {
	t.index++
	return t.currentVal()
}

func (t *TArrayIterator[T]) FirstVal() api.IZVal {
	t.index = 0
	return t.currentVal()
}
