package array

import (
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/errs"
	"github.com/lingdor/magicarray/internal"
	"reflect"
)

func ValueOfSlice[T any](val []T) MagicArray {
	if val == nil {
		return Make(false, false, 0)
	}
	return internal.TArray[T](val)
}

func ValueofStruct(val any) MagicArray {
	if val == nil {
		return Make(false, false, 0)
	}
	refVal := reflect.ValueOf(val)
	return internal.NewStructArray(val, refVal)
}
func ValueofStructs(list any) (MagicArray, error) {
	if list == nil {
		return Make(false, false, 0), nil
	}
	refVal := reflect.ValueOf(list)
	if refVal.Kind() != reflect.Slice {
		return nil, errs.TypeAssertError
	}
	return valueofStrucstLoad(list, refVal), nil
}
func ValueofMap[T any](m map[string]T) MagicArray {
	if m == nil {
		return Make(false, false, 0)
	}
	return internal.TMapArray[T](m)
}
func valueofStrucstLoad(list any, refVal reflect.Value) MagicArray {
	len := refVal.Len()
	arrs := make([]MagicArray, len)
	for i := 0; i < refVal.Len(); i++ {
		arrs[i] = ValueofStruct(refVal.Index(i).Interface())
	}
	return ValueOfSlice(arrs)
}

// NewArr MagicArray of array values
func NewArr[T any](vals ...T) MagicArray {
	return ValueOfSlice(vals)
}

// Valueof make a instance of MagicArray (no thread safe)
func Valueof(list any) (ret MagicArray, err error) {
	if list == nil {
		return nil, nil
		//return nil, errs.TypeAssertError
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
		if tmap, ok := list.(map[string]any); ok {
			return internal.TMapArray[any](tmap), nil
		} else if tmap, ok := list.(map[string]*any); ok {
			return internal.TMapArray[*any](tmap), nil
		}
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

func Clone(arr MagicArray) MagicArray {

	var mp = make(map[string]any, arr.Len())
	iter := arr.Iter()
	for k, v := iter.FirstKV(); v != nil; k, v = iter.NextKV() {
		mp[k.String()] = v.Interface()
	}
	return ValueofMap(mp)
}
