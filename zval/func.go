package zval

import (
	"database/sql/driver"
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/kind"
	"reflect"
	"time"
)

var ToMagicArr func(list any) (api.IMagicArray, error)

func NewZVal(val interface{}) api.IZVal {
	return NewZValOfReflect(val, nil)
}
func NewZValOfReflect(val any, refVal *reflect.Value) api.IZVal {

	if refVal == nil {
		rval := reflect.ValueOf(val)
		refVal = &rval
	}

	if refVal.Kind() == reflect.Ptr && refVal.IsNil() {
		return NewZValNil()
	}

NewZValOfReflect:
	switch zv := val.(type) {
	case api.IMagicArray:
		return NewZValOfKind(kind.MagicArray, val)
	case api.IZVal:
		return zv
	case driver.Valuer:
		if cv, err := zv.Value(); err == nil {
			if cv == nil {
				return NewZValNil()
			}
			return NewZValOfReflect(cv, nil)
		}
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
	case []byte:
		return NewZValOfKind(kind.Bytes, zv)
	}

	if uint8(refVal.Kind()) == kind.Pointer {
		refValRow := refVal.Elem()
		refVal = &refValRow
		val = refVal.Interface()
		goto NewZValOfReflect
	}
	return NewZValOfKind(uint8(refVal.Kind()), val)
}
func NewZValOfKind(kind uint8, val any) api.IZVal {
	return &ZValObj{
		kind: kind,
		val:  val,
	}
}

//func NewZValOfReflect(val reflect.Value) api.IZVal {
//	obj := val.Interface()
//	switch ins := obj.(type) {
//	case time.Time:
//		return NewZValOfKind(kind.Time, ins)
//	case api.IZVal:
//		return ins
//	case api.IMagicArray:
//		return NewZValOfKind(kind.MagicArray, obj)
//	}
//	return NewZValOfKind(uint8(val.Kind()), obj)
//}

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
