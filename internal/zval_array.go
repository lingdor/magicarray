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

func (m *ZValArray) Remove(key any) (api.WriteMagicArray, error) {
	if m.isKeys {
		var strKey string
		var ok bool
		if strKey, ok = key.(string); ok {
		} else if zval, ok := key.(api.IZVal); ok {
			strKey = zval.String()
		}
		if mapval, ok := m.mapVals[strKey]; ok {
			if m.isSorted {
				m.keys = append(m.keys[:mapval.index], m.keys[mapval.index+1:]...)
			}
			delete(m.mapVals, strKey)
		} else {
			return m, fmt.Errorf("%w map key:%s", errs.NoFundKey, strKey)
		}
		return m, nil
	}
	if intKey, ok := key.(int); ok {
		if m.Len() > intKey {
			m.listVals = append(m.listVals[:intKey], m.listVals[intKey+1:]...)
			return m, nil
		}
		return m, errs.OutOfArrayLength
	}

	return m, errs.TypeAssertError
}

func (m *ZValArray) Append(val any) api.WriteMagicArray {

	if m.isKeys {
		var i = 0
		for {
			i++
			if m.Get(i).IsNil() {
				m.Set(i, val)
			}
		}
	} else {
		m.listVals = append(m.listVals, zval.NewZVal(val))
	}
	return m
}

func (m *ZValArray) IsKeys() bool {
	return m.isKeys
}

func (m *ZValArray) Len() int {
	if m.isKeys {
		return len(m.mapVals)
	}
	return len(m.listVals)
}

func (m *ZValArray) Keys() api.IMagicArray {
	if !m.isKeys {
		keys := GenListKeys(m.Len())
		return TArray[int](keys)
	}
	if m.isSorted {
		return TArray[string](m.keys)
	}
	keys := make([]string, m.Len())
	i := -1
	for k, _ := range m.mapVals {
		i++
		keys[i] = k
	}
	return TArray[string](keys)
}

func (m *ZValArray) Values() api.IMagicArray {

	if !m.isKeys {
		return &ZValArray{
			isKeys:   false,
			listVals: m.listVals,
		}
	}
	vals := make([]api.IZVal, m.Len())
	iter := m.Iter()
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

func (m *ZValArray) Get(key interface{}) api.IZVal {
	if !m.isKeys {
		if index, ok := key.(int); ok {
			return m.listVals[index]
		}
		return zval.NewZValNil()
	}
	var zvalKey api.IZVal
	var ok bool
	if zvalKey, ok = key.(api.IZVal); !ok {
		zvalKey = zval.NewZVal(key)
	}
	if v, ok := m.mapVals[zvalKey.String()]; ok {
		return v.val
	}
	return zval.NewZValInvalid()
}

func (m *ZValArray) toMap() {
	l := len(m.listVals)
	//m.keys = make([]string, l)
	m.mapVals = make(map[string]ZValArrayMapVal, l)
	for i := 0; i < l; i++ {
		k := strconv.Itoa(i)
		m.mapVals[k] = ZValArrayMapVal{val: m.listVals[i], index: i}
	}
	m.isKeys = true
	m.listVals = nil
}

func (m *ZValArray) Set(key interface{}, val interface{}) api.WriteMagicArray {

	var zvalKey, zvalVal api.IZVal
	var ok bool
	if zvalKey, ok = key.(api.IZVal); !ok {
		zvalKey = zval.NewZVal(key)
	}
	if zvalVal, ok = val.(api.IZVal); !ok {
		zvalVal = zval.NewZVal(val)
	}
	if !m.isKeys {
		if intKey, ok := zvalKey.Int(); ok && intKey < m.Len() {
			m.listVals[intKey] = zvalVal
			return m
		}
		m.toMap()
	}

	if val, exists := m.mapVals[zvalKey.String()]; exists {
		m.mapVals[zvalKey.String()] = ZValArrayMapVal{val: zvalVal, index: val.index}
	} else {
		m.mapVals[zvalKey.String()] = ZValArrayMapVal{val: zvalVal, index: m.Len()}
		if m.isSorted {
			m.keys = append(m.keys, zvalKey.String())
		}
	}
	return m
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
func (m *ZValArray) MarshalJSON() ([]byte, error) {
	return JsonMarshal(m)
}
