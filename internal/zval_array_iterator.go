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
	reverse    bool
}

func (z *ZValArrayIterator) Index() int {
	return z.iteratePos
}

func (z *ZValArrayIterator) NextKV() (api.IZVal, api.IZVal) {
	if z.reverse == false {
		z.iteratePos++
	} else {
		z.iteratePos--
	}
	return z.currentKV()

}

func (z *ZValArrayIterator) FirstKV() (api.IZVal, api.IZVal) {
	if z.reverse == false {
		z.iteratePos = 0
	} else {
		z.iteratePos = z.arr.Len() - 1
	}
	return z.currentKV()
}

func (z *ZValArrayIterator) NextVal() api.IZVal {
	if z.reverse == false {
		z.iteratePos++
	} else {
		z.iteratePos--
	}
	return z.currentVal()

}

func (z *ZValArrayIterator) FirstVal() api.IZVal {
	if z.reverse == false {
		z.iteratePos = 0
	} else {
		z.iteratePos = z.arr.Len() - 1
	}
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
	if z.iteratePos < z.arr.Len() && z.iteratePos > -1 {
		if !z.arr.isKeys {
			return z.arr.listVals[z.iteratePos]
		}
		strKey := z.keys[z.iteratePos]
		return z.arr.mapVals[strKey]
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

func (z *ZValArray) RIter() api.Iterator {

	v := &ZValArrayIterator{
		arr:        z,
		iteratePos: -1,
		reverse:    false,
	}
	if z.isKeys {
		v.keys = z.Keys().(TArray[string])
	}
	return v
}
