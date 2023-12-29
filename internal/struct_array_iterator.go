package internal

import (
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
)

type StructArrayIterator struct {
	index int
	arr   *StructArray
	keys  []string
}

func (s *StructArrayIterator) Index() int {
	return s.index
}

func (i *StructArrayIterator) currentKV() (api.IZVal, api.IZVal) {
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
func (i *StructArrayIterator) NextKV() (api.IZVal, api.IZVal) {
	i.index++
	return i.currentKV()
}

func (i *StructArrayIterator) FirstKV() (api.IZVal, api.IZVal) {
	i.index = 0
	return i.currentKV()
}

func (i *StructArrayIterator) currentVal() api.IZVal {
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

func (i *StructArrayIterator) NextVal() api.IZVal {
	i.index++
	return i.currentVal()
}

func (i *StructArrayIterator) FirstVal() api.IZVal {
	i.index = 0
	return i.currentVal()
}
