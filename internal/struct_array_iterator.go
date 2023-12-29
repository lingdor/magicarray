package internal

import (
	"github.com/lingdor/magicarray/array/api"
	"github.com/lingdor/magicarray/array/kind"
	"github.com/lingdor/magicarray/array/zval"
)

type StructArrayIterator struct {
	index int
	arr   *StructArray
	keys  []string
}

func (i *StructArrayIterator) currentKV() (api.ZVal, api.ZVal) {
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
func (i *StructArrayIterator) NextKV() (api.ZVal, api.ZVal) {
	i.index++
	return i.currentKV()
}

func (i *StructArrayIterator) FirstKV() (api.ZVal, api.ZVal) {
	i.index = 0
	return i.currentKV()
}

func (i *StructArrayIterator) currentVal() api.ZVal {
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

func (i *StructArrayIterator) NextVal() api.ZVal {
	i.index++
	return i.currentVal()
}

func (i *StructArrayIterator) FirstVal() api.ZVal {
	i.index = 0
	return i.currentVal()
}
