package internal

import (
	"encoding/json"
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/zval"
	"reflect"
)

type MapArray struct {
	obj    any
	refVal reflect.Value
}

func (m *MapArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.obj)
}

func (s *MapArray) IsKeys() bool {
	return true
}

func NewMapArray(val any, refVal reflect.Value) *MapArray {

	return &MapArray{
		obj:    val,
		refVal: refVal,
	}
}

func (s *MapArray) Keys() api.IMagicArray {
	// todo can if map value type to generate difference array
	keys := s.genKeys()
	return TArray[any](keys)
}

func (s *MapArray) Values() api.IMagicArray {
	// todo can if map value type to generate difference array
	var vals = make([]any, 0, s.Len())
	iter := s.Iter()
	for val := iter.FirstVal(); val != nil; val = iter.NextVal() {
		vals = append(vals, val)
	}
	return TArray[any](vals)
}

func (s *MapArray) Len() int {
	return s.refVal.Len()
}

func (s MapArray) Get(key any) api.IZVal {
	var ok bool
	var strKey string
	if strKey, ok = key.(string); ok {
	} else if zvalKey, ok := key.(api.IZVal); ok {
		strKey = zvalKey.String()
	} else {
		strKey = zval.NewZVal(key).String()
	}

	keyVal := reflect.ValueOf(strKey)
	retVal := s.refVal.MapIndex(keyVal)
	return zval.NewZValOfReflect(retVal)
}
func (s *MapArray) genKeys() []any {

	keys := make([]any, s.Len())
	for index, kVal := range s.refVal.MapKeys() {
		keys[index] = kVal.Interface()
		index++
	}
	return keys
}

func (s *MapArray) Iter() api.Iterator {

	return &MapArrayIterator{
		arr:   s,
		index: 0,
		keys:  s.genKeys(),
	}
}
