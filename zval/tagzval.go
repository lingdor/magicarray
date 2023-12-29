package zval

import (
	"encoding/json"
	"github.com/lingdor/magicarray/api"
	"time"
)

type TagZVal struct {
	base api.ZVal
	tags map[string]string
}

func (t *TagZVal) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Interface())
}

func (t *TagZVal) Int() (int, bool) {
	return t.base.Int()
}

func (t *TagZVal) Uint() (uint, bool) {
	return t.base.Uint()
}

func (t *TagZVal) Compare(val api.ZVal) bool {
	return t.base.Compare(val)
}

func (t *TagZVal) Int64() (int64, bool) {
	return t.base.Int64()
}

func (t *TagZVal) Uint64() (uint64, bool) {
	return t.base.Uint64()
}

func (t *TagZVal) Int32() (int32, bool) {
	return t.base.Int32()
}

func (t *TagZVal) Int16() (int16, bool) {
	return t.base.Int16()
}

func (t *TagZVal) Int8() (int8, bool) {
	return t.base.Int8()
}

func (t *TagZVal) Uint32() (uint32, bool) {
	return t.base.Uint32()
}

func (t *TagZVal) Uint16() (uint16, bool) {
	return t.base.Uint16()
}

func (t *TagZVal) Uint8() (uint8, bool) {
	return t.base.Uint8()
}

func (t *TagZVal) String() string {
	return t.base.String()
}

func (t *TagZVal) ZVal() api.ZVal {
	return t.base.ZVal()
}

func (t *TagZVal) Interface() interface{} {
	return t.base.Interface()
}

func (t *TagZVal) Arr() (api.MagicArray, bool) {
	return t.base.Arr()
}

func (t *TagZVal) Kind() uint8 {
	return t.base.Kind()
}

func (t *TagZVal) IsEmpty() bool {
	return t.base.IsEmpty()
}

func (t *TagZVal) IsNil() bool {
	return t.base.IsNil()
}

func (t *TagZVal) IsSet() bool {
	return t.base.IsSet()
}

func (t *TagZVal) Float32() (float32, bool) {
	return t.base.Float32()
}

func (t *TagZVal) Float64() (float64, bool) {
	return t.base.Float64()
}

func (t *TagZVal) Time() (time.Time, bool) {
	return t.base.Time()
}

func (t *TagZVal) Bool() (bool, bool) {
	return t.base.Bool()
}

func (t *TagZVal) MustInt() int {
	return t.base.MustInt()
}

func (t *TagZVal) MustUint() uint {
	return t.base.MustUint()
}

func (t *TagZVal) MustInt32() int32 {
	return t.base.MustInt32()
}

func (t *TagZVal) MustInt16() int16 {
	return t.base.MustInt16()
}

func (t *TagZVal) MustInt8() int8 {
	return t.base.MustInt8()
}

func (t *TagZVal) MustUint32() uint32 {
	return t.base.MustUint32()
}

func (t *TagZVal) MustUint16() uint16 {
	return t.base.MustUint16()
}

func (t *TagZVal) MustUint8() uint8 {
	return t.base.MustUint8()
}

func (t *TagZVal) MustArr() api.MagicArray {
	return t.base.MustArr()
}

func (t *TagZVal) MustInt64() int64 {
	return t.base.MustInt64()
}

func (t *TagZVal) MustUint64() uint64 {
	return t.base.MustUint64()
}

func (t *TagZVal) MustFloat32() float32 {
	return t.base.MustFloat32()
}

func (t *TagZVal) MustFloat64() float64 {
	return t.base.MustFloat64()
}

func (t *TagZVal) MustTime() time.Time {
	return t.base.MustTime()
}

func (t *TagZVal) MustBool() bool {
	return t.base.MustBool()
}

func (t *TagZVal) SetTag(s string, val string) {
	t.tags[s] = val
}

func (t *TagZVal) GetTag(tag string) (string, bool) {
	if val, ok := t.tags[tag]; ok {
		return val, ok
	} else if tagval, ok := t.base.(api.ZValTag); ok {
		return tagval.GetTag(tag)
	} else {
		return "", false
	}
}

func NewTagZVal(zval api.ZVal) api.ZVal {
	return &TagZVal{
		base: zval,
		tags: make(map[string]string, 0),
	}
}
