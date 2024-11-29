package internal

import (
	"reflect"

	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/zval"
)

type MapArray struct {
	obj    any
	refVal reflect.Value
}

func (m *MapArray) IsKeys() bool {
	return true
}

func NewMapArray(val any, refVal reflect.Value) *MapArray {

	return &MapArray{
		obj:    val,
		refVal: refVal,
	}
}

func (m *MapArray) Keys() api.IMagicArray {
	// todo can if map value type to generate difference array
	keys := m.genKeys()
	return TArray[any](keys)
}

func (m *MapArray) Values() api.IMagicArray {
	// todo can if map value type to generate difference array
	var vals = make([]any, 0, m.Len())
	iter := m.Iter_()
	for val := iter.FirstVal(); val != nil; val = iter.NextVal() {
		vals = append(vals, val)
	}
	return TArray[any](vals)
}

func (m *MapArray) Len() int {
	return m.refVal.Len()
}

func (m *MapArray) Get(key any) api.IZVal {
	var ok bool
	var strKey string
	if strKey, ok = key.(string); ok {
	} else if zvalKey, ok := key.(api.IZVal); ok {
		strKey = zvalKey.String()
	} else {
		strKey = zval.NewZVal(key).String()
	}

	keyVal := reflect.ValueOf(strKey)
	retVal := m.refVal.MapIndex(keyVal)
	rawVal := retVal.Interface()
	if rawVal == nil {
		return zval.NewZValNil()
	}
	return zval.NewZVal(rawVal)
}
func (m *MapArray) genKeys() []any {

	keys := make([]any, m.Len())
	for index, kVal := range m.refVal.MapKeys() {
		keys[index] = kVal.Interface()
		index++
	}
	return keys
}

func (m *MapArray) Iter_() api.Iterator {

	return &MapArrayIterator{
		arr:   m,
		index: -1,
		keys:  m.genKeys(),
	}
}

func (m *MapArray) RIter_() api.Iterator {

	return &MapArrayIterator{
		arr:     m,
		index:   -1,
		keys:    m.genKeys(),
		reverse: true,
	}
}

func (m *MapArray) MarshalJSON() ([]byte, error) {
	return JsonMarshal(m)
}

func (m *MapArray) Iter() api.IterFunc {
	return func(yield func(api.IZVal, api.IZVal) bool) {
		iter := m.Iter_()
		for k, v := iter.FirstKV(); k != nil; k, v = iter.NextKV() {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (m *MapArray) IterRows() api.IterRowsFunc {
	return func(yield func(api.IZVal, api.IMagicArray) bool) {
		for k, v := range m.Iter() {
			if !yield(k, v.MustArr()) {
				return
			}
		}
	}
}
func (m *MapArray) Reverse() api.IMagicArray {
	return &ReverseMap{
		arr: m,
	}
}
