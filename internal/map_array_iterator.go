package internal

import (
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
)

type MapArrayIterator struct {
	index   int
	arr     *MapArray
	keys    []any
	reverse bool
}

func (m *MapArrayIterator) Index() int {
	return m.index
}

func (i *MapArrayIterator) currentKV() (api.IZVal, api.IZVal) {
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
func (i *MapArrayIterator) NextKV() (api.IZVal, api.IZVal) {
	if i.reverse == false {
		i.index++
	} else {
		i.index--
	}
	return i.currentKV()
}

func (i *MapArrayIterator) FirstKV() (api.IZVal, api.IZVal) {
	if i.reverse == false {
		i.index = 0
	} else {
		i.index = i.arr.Len() - 1
	}
	return i.currentKV()
}

func (i *MapArrayIterator) currentVal() api.IZVal {
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

func (i *MapArrayIterator) NextVal() api.IZVal {
	if i.reverse == false {
		i.index++
	} else {
		i.index--
	}
	return i.currentVal()
}

func (i *MapArrayIterator) FirstVal() api.IZVal {
	if i.reverse == false {
		i.index = 0
	} else {
		i.index = i.arr.Len() - 1
	}
	return i.currentVal()
}
