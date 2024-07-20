package internal

import (
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
)

type StructArrayIterator struct {
	index   int
	arr     *StructArray
	keys    []string
	reverse bool
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
	if i.reverse == false {
		i.index++
	} else {
		i.index--
	}
	return i.currentKV()
}

func (i *StructArrayIterator) FirstKV() (api.IZVal, api.IZVal) {
	if i.reverse == false {
		i.index = 0
	} else {
		i.index = i.arr.Len() - 1
	}
	return i.currentKV()
}

func (i *StructArrayIterator) currentVal() api.IZVal {
	if i.index < i.arr.Len() && i.index > -1 {
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
	if i.reverse == false {
		i.index++
	} else {
		i.index--
	}
	return i.currentVal()
}

func (i *StructArrayIterator) FirstVal() api.IZVal {
	if i.reverse == false {
		i.index = 0
	} else {
		i.index = i.arr.Len() - 1
	}
	return i.currentVal()
}
