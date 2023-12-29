package internal

import (
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
)

type MapArrayIterator struct {
	index int
	arr   *MapArray
	keys  []any
}

func (i *MapArrayIterator) currentKV() (api.ZVal, api.ZVal) {
	if i.index < i.arr.Len() {
		key := i.keys[i.index]
		if val := i.arr.Get(key); !val.IsNil() {
			return zval.NewZValOfKind(kind.String, key), val
		} else {
			return zval.NewZValOfKind(kind.String, key), zval.NewZValInvalid()
		}
	}
	return nil, nil
}
func (i *MapArrayIterator) NextKV() (api.ZVal, api.ZVal) {
	i.index++
	return i.currentKV()
}

func (i *MapArrayIterator) FirstKV() (api.ZVal, api.ZVal) {
	i.index = 0
	return i.currentKV()
}

func (i *MapArrayIterator) currentVal() api.ZVal {
	if i.index < i.arr.Len() {
		key := i.keys[i.index]
		if val := i.arr.Get(key); !val.IsNil() {
			return val
		} else {
			return zval.NewZValInvalid()
		}
	}
	return nil
}

func (i *MapArrayIterator) NextVal() api.ZVal {
	i.index++
	return i.currentVal()
}

func (i *MapArrayIterator) FirstVal() api.ZVal {
	i.index = 0
	return i.currentVal()
}
