package magicarray

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
		if arr, ok := v.Arr(); ok {
			col = append(col, arr.Get(key))
		}
	}
	return ValueOfSlice(col)
}
func Len(arr MagicArray) int {
	return arr.Len()
}
func Get(arr MagicArray, key interface{}) ZVal {
	return arr.Get(key)
}
func Keys(arr MagicArray) MagicArray {
	return arr.Keys()
}
func Values(arr MagicArray) MagicArray {
	return arr.Values()
}

// Pick Pick the keys and values to a new MagicArray for parameter keys order
func Pick(arr MagicArray, keys ...any) MagicArray {
	if arr.IsKeys() {
		var retKeys = make([]string, 0, len(keys))
		var retVals = make([]api.ZVal, 0, len(keys))
		for _, key := range keys {
			var strKey string
			var ok bool
			if strKey, ok = key.(string); !ok {
				strKey = zval.NewZVal(key).String()
			}
			retKeys = append(retKeys, strKey)
			retVals = append(retVals, arr.Get(key))
		}
		return internal.NewSortedArray(retKeys, retVals)
	} else {
		arr := internal.EmptyZValArray(false, false, len(keys))
		for _, key := range keys {
			var intKey int
			var ok bool
			if intKey, ok = key.(int); !ok {
				if intKey, ok = zval.NewZVal(key).Int(); !ok {
					continue
				}
			}
			if zvVal := arr.Get(intKey); zvVal.IsSet() {
				arr = Append(arr, zvVal)
			}
		}
		return arr
	}
}
