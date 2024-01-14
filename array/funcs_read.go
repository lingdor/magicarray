package array

import (
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/internal"
	"github.com/lingdor/magicarray/zval"
)

func Column(from MagicArray, key interface{}) MagicArray {

	var col = make([]ZVal, 0, from.Len())
	iter := from.Iter()
	var i = -1
	for v := iter.FirstVal(); v != nil; v = iter.NextVal() {
		i++
		if val, ok := v.Arr(); ok {
			col = append(col, val.Get(key))
		}
	}
	return ValueOfSlice(col)
}
func Len(marr MagicArray) int {
	return marr.Len()
}
func Get(marr MagicArray, key interface{}) ZVal {
	return marr.Get(key)
}
func Keys(marr MagicArray) MagicArray {
	return marr.Keys()
}
func Values(marr MagicArray) MagicArray {
	return marr.Values()
}

func MaxLen(marr MagicArray) int {
	max := 0
	iter := marr.Iter()
	for val := iter.FirstVal(); val != nil; val = iter.NextVal() {
		if val.IsNil() {
			continue
		}
		l := len(val.String())
		if l > max {
			max = l
		}
	}
	return max
}

// Pick Pick the keys and values to a new MagicArray for parameter keys order
func Pick(marr MagicArray, keys ...any) MagicArray {
	if marr.IsKeys() {
		var retKeys = make([]string, 0, len(keys))
		var retVals = make([]api.IZVal, 0, len(keys))
		for _, key := range keys {
			var strKey string
			var ok bool
			if strKey, ok = key.(string); !ok {
				strKey = zval.NewZVal(key).String()
			}
			retKeys = append(retKeys, strKey)
			retVals = append(retVals, marr.Get(key))
		}
		return internal.NewSortedArray(retKeys, retVals)
	} else {
		marr := internal.EmptyZValArray(false, false, len(keys))
		for _, key := range keys {
			var intKey int
			var ok bool
			if intKey, ok = key.(int); !ok {
				if intKey, ok = zval.NewZVal(key).Int(); !ok {
					continue
				}
			}
			if zvVal := marr.Get(intKey); zvVal.IsSet() {
				marr = Append(marr, zvVal)
			}
		}
		return marr
	}
}
