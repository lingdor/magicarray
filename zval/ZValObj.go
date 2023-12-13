package zval

import (
	"encoding/json"
	"fmt"
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/errs"
	"github.com/lingdor/magicarray/kind"
	"reflect"
	"strconv"
	"time"
)

type ZValObj struct {
	val  interface{}
	kind uint8
}

func (t *ZValObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Interface())
}

func (Z *ZValObj) IsSet() bool {
	return Z.kind != kind.Invalid
}

func (Z *ZValObj) Bool() (bool, bool) {
	if Z.kind == kind.Bool {
		v, ok := Z.val.(bool)
		return v, ok
	}
	if number, ok := Z.Int(); ok {
		return number == 1, true
	}
	return false, false
}

func (Z *ZValObj) MustBool() bool {
	if val, ok := Z.Bool(); ok {
		return val
	}
	panic(errs.TypeAssertError)
}

func (Z *ZValObj) Time() (time.Time, bool) {
	if Z.kind == kind.Time {
		v, ok := Z.val.(time.Time)
		return v, ok
	} else if Z.kind == kind.Int {
		return time.Unix(Z.MustInt64(), 0), true
	} else if Z.kind == kind.Int64 {
		return time.UnixMilli(Z.MustInt64()), true
	}
	return time.Time{}, false
}

func (Z *ZValObj) MustTime() time.Time {
	if val, ok := Z.Time(); ok {
		return val
	}
	panic(errs.TypeAssertError)
}

func (Z *ZValObj) Float32() (float32, bool) {
	if Z.kind == kind.Float32 {
		v, ok := Z.val.(float32)
		return v, ok
	}
	str := Z.String()
	if fval, err := strconv.ParseFloat(str, 32); err == nil {
		return float32(fval), false
	} else {
		return 0.0, false
	}

}

func (Z *ZValObj) Float64() (float64, bool) {
	if Z.kind == kind.Float64 {
		v, ok := Z.val.(float64)
		return v, ok
	}
	str := Z.String()
	if fval, err := strconv.ParseFloat(str, 64); err == nil {
		return fval, false
	} else {
		return 0.0, false
	}
}

func (Z *ZValObj) MustFloat32() float32 {
	if val, ok := Z.Float32(); ok {
		return val
	}
	panic(errs.TypeAssertError)
}

func (Z *ZValObj) MustFloat64() float64 {
	if val, ok := Z.Float64(); ok {
		return val
	}
	panic(errs.TypeAssertError)
}

func (Z *ZValObj) Uint() (uint, bool) {
	if Z.kind == kind.Uint {
		v, ok := Z.val.(uint)
		return v, ok
	}
	str := Z.String()

	if val, err := strconv.ParseUint(str, 10, strconv.IntSize); err == nil {
		return uint(val), false
	} else {
		return 0.0, false
	}
}

func (Z *ZValObj) Int64() (int64, bool) {
	if Z.kind == kind.Int64 {
		v, ok := Z.val.(int64)
		return v, ok
	}
	str := Z.String()

	if val, err := strconv.ParseInt(str, 10, 64); err == nil {
		return val, false
	} else {
		return 0.0, false
	}
}

func (Z *ZValObj) Uint64() (uint64, bool) {

	if Z.kind == kind.Uint64 {
		v, ok := Z.val.(uint64)
		return v, ok
	}
	str := Z.String()
	if val, err := strconv.ParseUint(str, 10, 64); err == nil {
		return val, false
	} else {
		return 0.0, false
	}

}

func (Z *ZValObj) MustUint() uint {

	if val, ok := Z.Uint(); ok {
		return val
	}
	panic(errs.TypeAssertError)
}

func (Z *ZValObj) MustInt64() int64 {
	if val, ok := Z.Int64(); ok {
		return val
	}
	panic(errs.TypeAssertError)
}

func (Z *ZValObj) MustUint64() uint64 {

	if val, ok := Z.Uint64(); ok {
		return val
	}
	panic(errs.TypeAssertError)
}

func (Z *ZValObj) MustInt() int {
	if val, ok := Z.Int(); ok {
		return val
	}
	panic(errs.TypeAssertError)
}

func (Z *ZValObj) MustInt32() int32 {
	if val, ok := Z.Int32(); ok {
		return val
	}
	panic(errs.TypeAssertError)
}

func (Z *ZValObj) MustInt16() int16 {
	if val, ok := Z.Int16(); ok {
		return val
	}
	panic(errs.TypeAssertError)
}

func (Z *ZValObj) MustInt8() int8 {
	if val, ok := Z.Int8(); ok {
		return val
	}
	panic(errs.TypeAssertError)
}

func (Z *ZValObj) MustUint32() uint32 {

	if val, ok := Z.Uint32(); ok {
		return val
	}
	panic(errs.TypeAssertError)
}

func (Z *ZValObj) MustUint16() uint16 {
	if val, ok := Z.Uint16(); ok {
		return val
	}
	panic(errs.TypeAssertError)
}

func (Z *ZValObj) MustUint8() uint8 {
	if val, ok := Z.Uint8(); ok {
		return val
	}
	panic(errs.TypeAssertError)
}

func (Z *ZValObj) MustArr() api.MagicArray {
	if val, ok := Z.Arr(); ok {
		return val
	}
	panic(errs.TypeAssertError)
}

func (Z *ZValObj) IsEmpty() bool {
	switch Z.kind {
	case kind.Invalid, kind.Nil:
		return true
	case kind.Int:
		return Z.MustInt() == 0
	case kind.Int64:
		return Z.MustInt64() == 0
	case kind.Array:
		return reflect.ValueOf(Z.val).Len() == 0
	case kind.MagicArray:
		return Z.MustArr().Len() == 0
	case kind.String:
		return len(Z.String()) == 0
	case kind.Slice:
		return reflect.ValueOf(Z.val).Len() == 0
	case kind.Int8:
		return Z.MustInt8() == 0
	case kind.Int16:
		return Z.MustInt16() == 0
	case kind.Int32:
		return Z.MustInt32() == 0
	case kind.Uint:
		return Z.MustUint() == 0
	case kind.Uint8:
		return Z.MustUint8() == 0
	case kind.Uint16:
		return Z.MustUint16() == 0
	case kind.Uint32:
		return Z.MustUint32() == 0
	case kind.Uint64:
		return Z.MustUint64() == 0
	case kind.Float32:
		return Z.MustFloat32() == 0
	case kind.Float64:
		return Z.MustInt8() == 0
	case kind.Bool:
		return Z.MustBool() == false
	default:
		return Z.val == nil
	}
}

func (Z *ZValObj) IsNil() bool {
	return Z.kind == kind.Nil || Z.kind == kind.Invalid
}

func (Z *ZValObj) Int() (int, bool) {
	str := ""
	if Z.kind == kind.Int {
		return Z.val.(int), true
	} else if Z.kind == kind.String {
		str = Z.val.(string)
	} else if Z.kind == kind.Float64 || Z.kind == kind.Float32 {
		//todo
		str = fmt.Sprintf("%.0f", Z.val)
	} else {
		str = fmt.Sprintf("%d", Z.val)
	}
	//todo have more well method
	if ret, err := strconv.Atoi(str); err != nil {
		return 0, false
	} else {
		return ret, true
	}
}

func (Z *ZValObj) Kind() uint8 {
	return Z.kind
}

func (Z *ZValObj) Compare(val api.ZVal) bool {
	if Z.Kind() == val.Kind() {
		return Z.val == val.Interface()
	}
	// todo give some other method
	return fmt.Sprintf("%v", Z.val) == fmt.Sprintf("%v", val.Interface())
}

func (Z *ZValObj) Int32() (int32, bool) {
	if Z.kind == kind.Int32 {
		v, ok := Z.val.(int32)
		return v, ok
	}
	str := Z.String()
	if val, err := strconv.ParseInt(str, 10, 32); err == nil {
		return int32(val), false
	} else {
		return 0, false
	}
}

func (Z *ZValObj) Int16() (int16, bool) {
	if Z.kind == kind.Int16 {
		v, ok := Z.val.(int16)
		return v, ok
	}
	str := Z.String()
	if val, err := strconv.ParseInt(str, 10, 16); err == nil {
		return int16(val), false
	} else {
		return 0, false
	}
}

func (Z *ZValObj) Int8() (int8, bool) {
	if Z.kind == kind.Int8 {
		v, ok := Z.val.(int8)
		return v, ok
	}
	str := Z.String()
	if val, err := strconv.ParseInt(str, 10, 8); err == nil {
		return int8(val), false
	} else {
		return 0, false
	}
}

func (Z *ZValObj) Uint32() (uint32, bool) {
	if Z.kind == kind.Uint32 {
		v, ok := Z.val.(uint32)
		return v, ok
	}
	str := Z.String()
	if val, err := strconv.ParseUint(str, 10, 32); err == nil {
		return uint32(val), false
	} else {
		return 0, false
	}
}

func (Z *ZValObj) Uint16() (uint16, bool) {
	if Z.kind == kind.Uint16 {
		v, ok := Z.val.(uint16)
		return v, ok
	}
	str := Z.String()
	if val, err := strconv.ParseUint(str, 10, 16); err == nil {
		return uint16(val), false
	} else {
		return 0, false
	}
}

func (Z *ZValObj) Uint8() (uint8, bool) {
	if Z.kind == kind.Uint8 {
		v, ok := Z.val.(uint8)
		return v, ok
	}
	str := Z.String()
	if val, err := strconv.ParseUint(str, 10, 8); err == nil {
		return uint8(val), false
	} else {
		return 0, false
	}
}

func (Z *ZValObj) String() string {
	if val, ok := Z.val.(string); ok {
		return val
	}
	return fmt.Sprintf("%v", Z.val)
}

func (Z *ZValObj) ZVal() api.ZVal {
	return Z
}

func (Z *ZValObj) Interface() any {
	return Z.val
}

func (Z *ZValObj) Arr() (api.MagicArray, bool) {
	if val, ok := Z.val.(api.MagicArray); ok {
		return val, true
	}
	if arr, err := ToMagicArr(Z.val); err == nil {
		return arr, true
	} else {
		return nil, false
	}
}
