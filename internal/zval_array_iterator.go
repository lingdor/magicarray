package internal

import (
	"github.com/lingdor/magicarray/array/api"
	"github.com/lingdor/magicarray/array/kind"
	"github.com/lingdor/magicarray/array/zval"
)

type ZValArrayIterator struct {
	iteratePos int
	keys       []string
	arr        *ZValArray
}

func (z *ZValArrayIterator) NextKV() (api.ZVal, api.ZVal) {
	z.iteratePos++
	return z.currentKV()

}

func (z *ZValArrayIterator) FirstKV() (api.ZVal, api.ZVal) {
	z.iteratePos = 0
	return z.currentKV()
}

func (z *ZValArrayIterator) NextVal() api.ZVal {
	z.iteratePos++
	return z.currentVal()

}

func (z *ZValArrayIterator) FirstVal() api.ZVal {
	z.iteratePos = 0
	return z.currentVal()
}

func (z *ZValArrayIterator) currentKV() (api.ZVal, api.ZVal) {
	if z.iteratePos < z.arr.Len() {
		if !z.arr.isKeys {
			return zval.NewZValOfKind(kind.Int, z.iteratePos), z.arr.listVals[z.iteratePos]
		}
		strKey := z.keys[z.iteratePos]
		return zval.NewZValOfKind(kind.String, strKey), z.arr.Get(strKey)
	} else {
		return nil, nil
	}
}

func (z *ZValArrayIterator) currentVal() api.ZVal {
	if z.iteratePos < z.arr.Len() {
		if !z.arr.isKeys {
			return z.arr.listVals[z.iteratePos]
		}
		strKey := z.keys[z.iteratePos]
		return z.arr.mapVals[strKey].val
	} else {
		return nil
	}
}

func (m *ZValArray) Iter() api.Iterator {

	v := &ZValArrayIterator{
		arr:        m,
		iteratePos: 0,
	}
	if m.isKeys {
		v.keys = m.Keys().(TArray[string])
	}
	return v
}
