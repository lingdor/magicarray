package array

import (
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/errs"
	"github.com/lingdor/magicarray/internal"
	"reflect"
)

func ValueOfSlice[T any](val []T) MagicArray {
	return internal.TArray[T](val)
}

func ValueofStruct(val any) MagicArray {
	refVal := reflect.ValueOf(val)
	return internal.NewStructArray(val, refVal)
}
func ValueofStructs(list any) (MagicArray, error) {
	refVal := reflect.ValueOf(list)
	if refVal.Kind() != reflect.Slice || refVal.Elem().Kind() != reflect.Struct {
		return nil, errs.TypeAssertError
	}
	return valueofStrucstLoad(list, refVal), nil

}
func valueofStrucstLoad(list any, refVal reflect.Value) MagicArray {
	len := refVal.Len()
	arrs := make([]MagicArray, len)
	for i := 0; i < refVal.Len(); i++ {
		arrs[i] = ValueofStruct(refVal.Index(i).Interface())
	}
	return ValueOfSlice(arrs)
}

// Valueof make a instance of MagicArray (no thread safe)
func Valueof(list any) (ret MagicArray, err error) {
	if list == nil {
		return nil, errs.TypeAssertError
	}
	if arr, ok := list.(MagicArray); ok {
		return arr, nil
	}
	refVal := reflect.ValueOf(list)
	kind := refVal.Kind()
	refType := refVal.Type()
	if kind == reflect.Slice {
		switch refType.Elem().Kind() {
		case reflect.Struct:
			return valueofStrucstLoad(list, refVal), nil
		case reflect.Int:
			if ints, ok := list.([]int); !ok {
				return nil, errs.TypeAssertError
			} else {
				return ValueOfSlice(ints), nil
			}
		case reflect.String:
			if strs, ok := list.([]string); !ok {
				return nil, errs.TypeAssertError
			} else {
				return ValueOfSlice(strs), nil
			}
		case reflect.Pointer:
			val := refVal.Elem().Interface()
			if val == list {
				return nil, errs.LinksDeepOutError
			}
			return Valueof(val)
		default:
			if ss, ok := list.([]MagicArray); ok {
				return ValueOfSlice(ss), nil
			} else if ss, ok := list.([]api.IMagicArray); ok {
				return ValueOfSlice(ss), nil
			}
			objs := internal.SlicetoAnyList(refVal)
			return ValueOfSlice(objs), nil
		}
	} else if kind == reflect.Map {

		return internal.NewMapArray(list, refVal), nil
	} else if kind == reflect.Struct {
		return ValueofStruct(list), nil
	} else {
		return nil, errs.TypeAssertError
	}
}
func MustValueof(list any) MagicArray {
	if arr, err := Valueof(list); err == nil {
		return arr
	} else {
		panic(err)
	}
}
func Make(isKeys, isSort bool, cap int) MagicArray {
	return internal.EmptyZValArray(isKeys, isSort, cap)
}
