package internal

import (
	"encoding/json"
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/zval"
	"reflect"
)

type StructArray struct {
	obj    any
	refVal reflect.Value
}

func (m *StructArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.obj)
}

func (s *StructArray) IsKeys() bool {
	return true
}

func NewStructArray(val any, refVal reflect.Value) *StructArray {

	return &StructArray{
		obj:    val,
		refVal: refVal,
	}
}

func (s *StructArray) Keys() api.MagicArray {
	keys := s.genKeys()
	return TArray[string](keys)
}

func (s *StructArray) Values() api.MagicArray {
	var vals = make([]any, 0, s.Len())
	iter := s.Iter()
	for val := iter.FirstVal(); val != nil; val = iter.NextVal() {
		vals = append(vals, val)
	}
	return TArray[any](vals)
}

func (s *StructArray) Len() int {
	return s.refVal.NumField()
}

func (s StructArray) Get(key any) api.ZVal {
	var ok bool
	var strKey string
	if strKey, ok = key.(string); !ok {
		return zval.NewZValInvalid()
	} else if zvalKey, ok := key.(api.ZVal); ok {
		strKey = zvalKey.String()
	} else {
		strKey = zval.NewZVal(key).String()
	}

	if typeField, ok := s.refVal.Type().FieldByName(strKey); ok {
		refVal := s.refVal.FieldByName(typeField.Name)
		return zval.NewStructTagZVal(zval.NewZValOfReflect(refVal), typeField.Tag)
	}
	return zval.NewZValInvalid()
}
func (s *StructArray) genKeys() []string {

	t := s.refVal.Type()
	keys := make([]string, s.Len())
	for i := 0; i < s.Len(); i++ {
		keys[i] = t.Field(i).Name
	}
	return keys
}

func (s *StructArray) Iter() api.Iterator {

	return &StructArrayIterator{
		arr:   s,
		index: 0,
		keys:  s.genKeys(),
	}
}
