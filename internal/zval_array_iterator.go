package internal

import (
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
)

type ZValArrayIterator struct {
	iteratePos int
	keys       []string
	arr        *ZValArray
}

func (z *ZValArrayIterator) Index() int {
	return z.iteratePos
}

func (z *ZValArrayIterator) NextKV() (api.IZVal, api.IZVal) {
	z.iteratePos++
	return z.currentKV()

}

func (z *ZValArrayIterator) FirstKV() (api.IZVal, api.IZVal) {
	z.iteratePos = 0
	return z.currentKV()
}

func (z *ZValArrayIterator) NextVal() api.IZVal {
	z.iteratePos++
	return z.currentVal()

}

func (z *ZValArrayIterator) FirstVal() api.IZVal {
	z.iteratePos = 0
	return z.currentVal()
}

func (z *ZValArrayIterator) currentKV() (api.IZVal, api.IZVal) {
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

func (z *ZValArrayIterator) currentVal() api.IZVal {
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

func (z *ZValArray) Iter() api.Iterator {

	v := &ZValArrayIterator{
		arr:        z,
		iteratePos: -1,
	}
	if z.isKeys {
		v.keys = z.Keys().(TArray[string])
	}
	return v
}
