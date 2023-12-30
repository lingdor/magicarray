package internal

import (
	"fmt"
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/errs"
	"github.com/lingdor/magicarray/zval"
	"strconv"
)

type ZValArray struct {
	keys     []string
	isKeys   bool
	isSorted bool
	mapVals  map[string]ZValArrayMapVal
	listVals []api.IZVal
}

type ZValArrayMapVal struct {
	val   api.IZVal
	index int
}

func (z *ZValArray) Remove(key any) (api.WriteMagicArray, error) {
	if z.isKeys {
		var strKey string
		var ok bool
		if strKey, ok = key.(string); ok {
		} else if zval, ok := key.(api.IZVal); ok {
			strKey = zval.String()
		}
		if mapval, ok := z.mapVals[strKey]; ok {
			if z.isSorted {
				z.keys = append(z.keys[:mapval.index], z.keys[mapval.index+1:]...)
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
		return v.val
	}
	return zval.NewZValInvalid()
}

func (z *ZValArray) toMap() {
	l := len(z.listVals)
	//z.keys = make([]string, l)
	z.mapVals = make(map[string]ZValArrayMapVal, l)
	for i := 0; i < l; i++ {
		k := strconv.Itoa(i)
		z.mapVals[k] = ZValArrayMapVal{val: z.listVals[i], index: i}
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

	if val, exists := z.mapVals[zvalKey.String()]; exists {
		z.mapVals[zvalKey.String()] = ZValArrayMapVal{val: zvalVal, index: val.index}
	} else {
		z.mapVals[zvalKey.String()] = ZValArrayMapVal{val: zvalVal, index: z.Len()}
		if z.isSorted {
			z.keys = append(z.keys, zvalKey.String())
		}
	}
	return z
}
func EmptyZValArray(isKeys, isSort bool, cap int) api.IMagicArray {
	return &ZValArray{
		keys:     make([]string, 0, cap),
		isKeys:   isKeys,
		isSorted: isSort,
		mapVals:  make(map[string]ZValArrayMapVal, cap),
	}
}
func NewSortedArray(keys []string, vals []api.IZVal) api.IMagicArray {
	mapVals := make(map[string]ZValArrayMapVal, len(keys))
	for i := 0; i < len(vals); i++ {
		mapVals[keys[i]] = ZValArrayMapVal{
			val:   vals[i],
			index: i,
		}
	}
	return &ZValArray{
		keys:     keys,
		isSorted: true,
		isKeys:   true,
		mapVals:  mapVals,
	}
}
func (z *ZValArray) MarshalJSON() ([]byte, error) {
	return JsonMarshal(z)
}
