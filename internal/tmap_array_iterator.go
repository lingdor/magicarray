package internal

import (
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
)

type TMapArrayIterator[T any] struct {
	index int
	arr   TMapArray[T]
	keys  []string
}

func (t *TMapArrayIterator[T]) Index() int {
	return t.index
}

func (t *TMapArrayIterator[T]) currentKV() (api.IZVal, api.IZVal) {
	if t.index < t.arr.Len() {
		key := t.keys[t.index]
		if val := t.arr.Get(key); !val.IsNil() {
			return zval.NewZValOfKind(kind.String, key), val
		} else {
			return zval.NewZValOfKind(kind.String, key), zval.NewZValInvalid()
		}
	}
	return nil, nil
}
func (t *TMapArrayIterator[T]) NextKV() (api.IZVal, api.IZVal) {
	t.index++
	return t.currentKV()
}

func (t *TMapArrayIterator[T]) FirstKV() (api.IZVal, api.IZVal) {
	t.index = 0
	return t.currentKV()
}

func (t *TMapArrayIterator[T]) currentVal() api.IZVal {
	if t.index < t.arr.Len() {
		key := t.keys[t.index]
		if val := t.arr.Get(key); !val.IsNil() {
			return val
		} else {
			return zval.NewZValInvalid()
		}
	}
	return nil
}

func (t *TMapArrayIterator[T]) NextVal() api.IZVal {
	t.index++
	return t.currentVal()
}

func (t *TMapArrayIterator[T]) FirstVal() api.IZVal {
	t.index = 0
	return t.currentVal()
}
