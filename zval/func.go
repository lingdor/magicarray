package zval

import (
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/kind"
	"reflect"
	"time"
)

var ToMagicArr func(list any) (api.IMagicArray, error)

func NewZVal(val interface{}) api.IZVal {
	switch zv := val.(type) {
	case api.IMagicArray:
		return NewZValOfKind(kind.MagicArray, val)
	case api.IZVal:
		return zv
	case string:
		return NewZValOfKind(kind.String, val)
	case int:
		return NewZValOfKind(kind.Int, val)
	case uint:
		return NewZValOfKind(kind.Uint, val)
	case int64:
		return NewZValOfKind(kind.Int64, val)
	case uint64:
		return NewZValOfKind(kind.Uint64, val)
	case int8:
		return NewZValOfKind(kind.Int8, val)
	case uint8:
		return NewZValOfKind(kind.Uint8, val)
	case int16:
		return NewZValOfKind(kind.Int16, val)
	case uint16:
		return NewZValOfKind(kind.Uint16, val)
	case int32:
		return NewZValOfKind(kind.Int32, val)
	case uint32:
		return NewZValOfKind(kind.Uint32, val)
	case float32:
		return NewZValOfKind(kind.Float32, val)
	case float64:
		return NewZValOfKind(kind.Float64, val)
	case time.Time:
		return NewZValOfKind(kind.Time, val)
	case bool:
		return NewZValOfKind(kind.Bool, val)
	}
	refVal := reflect.ValueOf(val)
	return NewZValOfReflect(refVal)
}
func NewZValOfKind(kind uint8, val any) api.IZVal {
	return &ZValObj{
		kind: kind,
		val:  val,
	}
}
func NewZValOfReflect(val reflect.Value) api.IZVal {

	obj := val.Interface()
	if _, ok := obj.(api.IMagicArray); ok {
		return NewZValOfKind(kind.MagicArray, obj)
	} else if v, ok := obj.(api.IZVal); ok {
		return v
	}
	return NewZValOfKind(uint8(val.Kind()), obj)
}

func NewZValInvalid() api.IZVal {

	return &ZValObj{
		kind: kind.Invalid,
		val:  nil,
	}
}

func NewZValNil() api.IZVal {

	return &ZValObj{
		kind: kind.Interface,
		val:  nil,
	}
}
