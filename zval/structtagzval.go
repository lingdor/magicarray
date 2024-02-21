package zval

import (
	"encoding/json"
	"github.com/lingdor/magicarray/api"
	"reflect"
	"time"
)

type StructTagZVal struct {
	base     api.IZVal
	fieldTag reflect.StructTag
}

func (t *StructTagZVal) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Interface())
}

func (t *StructTagZVal) Int() (int, bool) {
	return t.base.Int()
}

func (t *StructTagZVal) Uint() (uint, bool) {
	return t.base.Uint()
}

func (t *StructTagZVal) Compare(val api.IZVal) bool {
	return t.base.Compare(val)
}

func (t *StructTagZVal) Int64() (int64, bool) {
	return t.base.Int64()
}

func (t *StructTagZVal) Uint64() (uint64, bool) {
	return t.base.Uint64()
}

func (t *StructTagZVal) Int32() (int32, bool) {
	return t.base.Int32()
}

func (t *StructTagZVal) Int16() (int16, bool) {
	return t.base.Int16()
}

func (t *StructTagZVal) Int8() (int8, bool) {
	return t.base.Int8()
}

func (t *StructTagZVal) Uint32() (uint32, bool) {
	return t.base.Uint32()
}

func (t *StructTagZVal) Uint16() (uint16, bool) {
	return t.base.Uint16()
}

func (t *StructTagZVal) Uint8() (uint8, bool) {
	return t.base.Uint8()
}

func (t *StructTagZVal) String() string {
	return t.base.String()
}

func (t *StructTagZVal) ZVal() api.IZVal {
	return t.base.ZVal()
}

func (t *StructTagZVal) Interface() interface{} {
	return t.base.Interface()
}

func (t *StructTagZVal) Arr() (api.IMagicArray, bool) {
	return t.base.Arr()
}

func (t *StructTagZVal) Kind() uint8 {
	return t.base.Kind()
}

func (t *StructTagZVal) IsEmpty() bool {
	return t.base.IsEmpty()
}

func (t *StructTagZVal) IsNil() bool {
	return t.base.IsNil()
}

func (t *StructTagZVal) IsSet() bool {
	return t.base.IsSet()
}

func (t *StructTagZVal) Float32() (float32, bool) {
	return t.base.Float32()
}

func (t *StructTagZVal) Float64() (float64, bool) {
	return t.base.Float64()
}

func (t *StructTagZVal) Time() (time.Time, bool) {
	return t.base.Time()
}

func (t *StructTagZVal) Bool() (bool, bool) {
	return t.base.Bool()
}

func (t *StructTagZVal) MustInt() int {
	return t.base.MustInt()
}

func (t *StructTagZVal) MustUint() uint {
	return t.base.MustUint()
}

func (t *StructTagZVal) MustInt32() int32 {
	return t.base.MustInt32()
}

func (t *StructTagZVal) MustInt16() int16 {
	return t.base.MustInt16()
}

func (t *StructTagZVal) MustInt8() int8 {
	return t.base.MustInt8()
}

func (t *StructTagZVal) MustUint32() uint32 {
	return t.base.MustUint32()
}

func (t *StructTagZVal) MustUint16() uint16 {
	return t.base.MustUint16()
}

func (t *StructTagZVal) MustUint8() uint8 {
	return t.base.MustUint8()
}

func (t *StructTagZVal) MustArr() api.IMagicArray {
	return t.base.MustArr()
}

func (t *StructTagZVal) MustInt64() int64 {
	return t.base.MustInt64()
}

func (t *StructTagZVal) MustUint64() uint64 {
	return t.base.MustUint64()
}

func (t *StructTagZVal) MustFloat32() float32 {
	return t.base.MustFloat32()
}

func (t *StructTagZVal) MustFloat64() float64 {
	return t.base.MustFloat64()
}

func (t *StructTagZVal) MustTime() time.Time {
	return t.base.MustTime()
}

func (t *StructTagZVal) MustBool() bool {
	return t.base.MustBool()
}

func (f *StructTagZVal) GetTag(tag string) (string, bool) {
	return f.fieldTag.Lookup(tag)
}

func (f *StructTagZVal) Bytes() []byte {
	return f.base.Bytes()
}

func NewStructTagZVal(zval api.IZVal, field reflect.StructTag) api.IZVal {
	return &StructTagZVal{
		base:     zval,
		fieldTag: field,
	}
}
