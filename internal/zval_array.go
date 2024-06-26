package internal

import (
	"fmt"
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/errs"
	"github.com/lingdor/magicarray/zval"
	"strconv"
)

type ZValArray struct {
	keys      []string
	KeysIndex map[string]int
	isKeys    bool
	isSorted  bool
	mapVals   map[string]api.IZVal
	listVals  []api.IZVal
}

type ZValArrayMapVal struct {
	val   api.IZVal
	index int
}

func (z *ZValArray) tidyKeysIndex() {
	z.KeysIndex = make(map[string]int, z.Len())
	for i, k := range z.keys {
		z.KeysIndex[k] = i
	}
}

func (z *ZValArray) Remove(key any) (api.WriteMagicArray, error) {
	if z.isKeys {
		var strKey string
		var ok bool
		if strKey, ok = key.(string); ok {
		} else if zval, ok := key.(api.IZVal); ok {
			strKey = zval.String()
		}
		if _, ok := z.mapVals[strKey]; ok {
			if z.isSorted {
				idx := z.KeysIndex[strKey]
				z.keys = append(z.keys[:idx], z.keys[idx+1:]...)
				z.tidyKeysIndex()
			}
			delete(z.mapVals, strKey)
		} else {
			return z, fmt.Errorf("%w map key:%s", errs.NoFundKey, strKey)
		}
		return z, nil
	}
	if intKey, ok := key.(int); ok {
		if z.Len() > intKey {
			z.listVals = append(z.listVals[:intKey], z.listVals[intKey+1:]...)
			return z, nil
		}
		return z, errs.OutOfArrayLength
	}

	return z, errs.TypeAssertError
}

func (z *ZValArray) Append(val any) api.WriteMagicArray {

	if z.isKeys {
		var i = 0
		for {
			i++
			if z.Get(i).IsNil() {
				z.Set(i, val)
				return z
			}
		}
	} else {
		z.listVals = append(z.listVals, zval.NewZVal(val))
	}
	return z
}

func (z *ZValArray) IsKeys() bool {
	return z.isKeys
}

func (z *ZValArray) Len() int {
	if z.isKeys {
		return len(z.mapVals)
	}
	return len(z.listVals)
}

func (z *ZValArray) Keys() api.IMagicArray {
	if !z.isKeys {
		keys := GenListKeys(z.Len())
		return TArray[int](keys)
	}
	if z.isSorted {
		return TArray[string](z.keys)
	}
	keys := make([]string, z.Len())
	i := -1
	for k, _ := range z.mapVals {
		i++
		keys[i] = k
	}
	return TArray[string](keys)
}

func (z *ZValArray) Values() api.IMagicArray {

	if !z.isKeys {
		return &ZValArray{
			isKeys:   false,
			listVals: z.listVals,
		}
	}
	vals := make([]api.IZVal, z.Len())
	iter := z.Iter()
	var i = -1
	for v := iter.FirstVal(); v != nil; v = iter.NextVal() {
		i++
		vals[i] = v
	}
	return &ZValArray{
		isKeys:   false,
		listVals: vals,
	}
}

func (z *ZValArray) Get(key interface{}) api.IZVal {
	if !z.isKeys {
		if index, ok := key.(int); ok {
			return z.listVals[index]
		}
		return zval.NewZValNil()
	}
	var zvalKey api.IZVal
	var ok bool
	if zvalKey, ok = key.(api.IZVal); !ok {
		zvalKey = zval.NewZVal(key)
	}
	if v, ok := z.mapVals[zvalKey.String()]; ok {
		return v
	}
	return zval.NewZValInvalid()
}

func (z *ZValArray) toMap() {
	l := len(z.listVals)
	//z.keys = make([]string, l)
	z.mapVals = make(map[string]api.IZVal, l)
	for i := 0; i < l; i++ {
		k := strconv.Itoa(i)
		z.mapVals[k] = z.listVals[i]
	}
	z.isKeys = true
	z.listVals = nil
}

func (z *ZValArray) Set(key interface{}, val interface{}) api.WriteMagicArray {

	var zvalKey, zvalVal api.IZVal
	var ok bool
	if zvalKey, ok = key.(api.IZVal); !ok {
		zvalKey = zval.NewZVal(key)
	}
	if zvalVal, ok = val.(api.IZVal); !ok {
		zvalVal = zval.NewZVal(val)
	}
	if !z.isKeys {
		if intKey, ok := zvalKey.Int(); ok && intKey < z.Len() {
			z.listVals[intKey] = zvalVal
			return z
		}
		z.toMap()
	}

	if _, exists := z.mapVals[zvalKey.String()]; exists {
		z.mapVals[zvalKey.String()] = zvalVal
	} else {
		z.mapVals[zvalKey.String()] = zvalVal
		if z.isSorted {
			z.keys = append(z.keys, zvalKey.String())
			z.KeysIndex[zvalKey.String()] = len(z.keys) - 1
		}
	}
	return z
}
func EmptyZValArray(isKeys, isSort bool, cap int) api.IMagicArray {
	return &ZValArray{
		keys:      make([]string, 0, cap),
		KeysIndex: make(map[string]int),
		isKeys:    isKeys,
		isSorted:  isSort,
		mapVals:   make(map[string]api.IZVal, cap),
	}
}
func NewSortedArray(keys []string, vals []api.IZVal) api.IMagicArray {
	if len(keys) != len(vals) {
		return nil
	}

	mapVals := make(map[string]api.IZVal, len(keys))
	keyIndexMap := make(map[string]int, len(keys))
	for i := 0; i < len(vals); i++ {
		mapVals[keys[i]] = vals[i]
		keyIndexMap[keys[i]] = i
	}

	return &ZValArray{
		keys:      keys,
		isSorted:  true,
		isKeys:    true,
		mapVals:   mapVals,
		KeysIndex: keyIndexMap,
	}
}
func (z *ZValArray) MarshalJSON() ([]byte, error) {
	return JsonMarshal(z)
}
