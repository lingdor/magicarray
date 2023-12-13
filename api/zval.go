package api

import (
	"time"
)

type ZVal interface {

	// Compare less=1 equal = 2 large =3 faild = 4
	Int() (int, bool)
	Uint() (uint, bool)
	Compare(val ZVal) bool
	Int64() (int64, bool)
	Uint64() (uint64, bool)
	Int32() (int32, bool)
	Int16() (int16, bool)
	Int8() (int8, bool)
	Uint32() (uint32, bool)
	Uint16() (uint16, bool)
	Uint8() (uint8, bool)
	String() string
	ZVal() ZVal
	Interface() interface{}
	Arr() (MagicArray, bool)
	Kind() uint8
	IsEmpty() bool
	IsNil() bool
	//IsSet is different with php, nil whil return true
	IsSet() bool
	Float32() (float32, bool)
	Float64() (float64, bool)
	Time() (time.Time, bool)
	Bool() (bool, bool)

	MustInt() int
	MustUint() uint
	MustInt32() int32
	MustInt16() int16
	MustInt8() int8
	MustUint32() uint32
	MustUint16() uint16
	MustUint8() uint8
	MustArr() MagicArray
	MustInt64() int64
	MustUint64() uint64
	MustFloat32() float32
	MustFloat64() float64
	MustTime() time.Time
	MustBool() bool
}
