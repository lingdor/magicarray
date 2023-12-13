package internal

import (
	"github.com/lingdor/magicarray/api"
	"reflect"
)

func GenListKeys(len int) []int {
	var keys = make([]int, len)
	for i := 0; i < len; i++ {
		keys[i] = i
	}
	return keys
}

func SlicetoAnyList(refVal reflect.Value) []any {
	len := refVal.Len()
	ret := make([]any, 0, len)
	for i := 0; i < len; i++ {
		ret = append(ret, refVal.Index(i).Interface())
	}
	return ret
}
func newTArray[T any](listVal []T) api.MagicArray {
	return TArray[T](listVal)
}

//
//func NamingJsonHump(arr api.MagicArray) api.MagicArray {
//
//	if arr == nil {
//		return &ZValArray{isKeys: false, listVals: []api.ZVal{}}
//	}
//	if arr.IsKeys() {
//		var newKeys = make([]string, arr.Len())
//		var newVals = make(map[string]ZValArrayMapVal, arr.Len())
//		var refType api.RefType
//		var refTypeOK bool
//		refType, refTypeOK = arr.(api.RefType)
//		var iter = arr.Iter()
//		var i = -1
//		for k, v := iter.FirstKV(); k != nil; k, v = iter.NextKV() {
//			i++
//			ismatch := false
//			if refTypeOK {
//				tt := refType.GetRefType(k.String())
//				if jsonName, ok := tt.Tag.Lookup("json"); ok {
//					newKeys[i] = jsonName
//					ismatch = true
//				}
//			}
//			if !ismatch && len(k.String()) > 0 {
//				runes := []rune(k.String())
//				runes[0] = unicode.ToLower(runes[0])
//				newKeys[i] = string(runes)
//			}
//			newVals[newKeys[i]] = ZValArrayMapVal{val: v, index: i}
//		}
//		return &ZValArray{
//			keys:     newKeys,
//			isSorted: true,
//			isKeys:   true,
//			mapVals:  newVals,
//		}
//	}
//
//	vals := make([]api.ZVal, arr.Len())
//	for i := 0; i < arr.Len(); i++ {
//		v := arr.Get(i)
//		if child, ok := v.Arr(); ok {
//			vals[i] = zval.NewZValOfKind(kind.MagicArray, NamingJsonHump(child))
//		} else {
//			vals[i] = v
//		}
//	}
//	return TArray[api.ZVal](vals)
//
//}
