package internal

import (
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
)

type TMapArrayIterator[T any] struct {
	index   int
	arr     TMapArray[T]
	keys    []string
	reverse bool
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
	if t.reverse == false {
		t.index++
	} else {
		t.index--
	}
	return t.currentKV()
}

func (t *TMapArrayIterator[T]) FirstKV() (api.IZVal, api.IZVal) {
	if t.reverse == false {
		t.index = 0
	} else {
		t.index = t.arr.Len() - 1
	}
	return t.currentKV()
}

func (t *TMapArrayIterator[T]) currentVal() api.IZVal {
	if t.index < t.arr.Len() && t.index > -1 {
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
	if t.reverse == false {
		t.index++
	} else {
		t.index--
	}
	return t.currentVal()
}

func (t *TMapArrayIterator[T]) FirstVal() api.IZVal {
	if t.reverse == false {
		t.index = 0
	} else {
		t.index = t.arr.Len() - 1
	}
	return t.currentVal()
}
