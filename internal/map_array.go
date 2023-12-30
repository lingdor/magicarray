package internal

import (
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/zval"
	"reflect"
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
	iter := m.Iter()
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

func (m *MapArray) Iter() api.Iterator {

	return &MapArrayIterator{
		arr:   m,
		index: 0,
		keys:  m.genKeys(),
	}
}

func (m *MapArray) MarshalJSON() ([]byte, error) {
	return JsonMarshal(m)
}
