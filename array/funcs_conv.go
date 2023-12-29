package array

import (
	"github.com/lingdor/magicarray/internal"
)

func ToWriter(marr MagicArray) WriteMagicArray {
	var setter WriteMagicArray
	var ok bool
	if setter, ok = marr.(WriteMagicArray); ok {
		return setter
	}
	setter = Make(marr.IsKeys(), true, marr.Len()).(WriteMagicArray)
	iter := marr.Iter()
	for k, v := iter.FirstKV(); k != nil; k, v = iter.NextKV() {
		if marr.IsKeys() {
			setter = setter.Set(k, v)
		} else {
			setter = setter.Append(v)
		}
	}
	return setter
}

func ToStringList(array MagicArray) []string {
	if strs, ok := array.(internal.TArray[string]); ok {
		return []string(strs)
	}
	strs := make([]string, 0, array.Len())
	iter := array.Iter()
	for val := iter.FirstVal(); val != nil; val = iter.NextVal() {
		strs = append(strs, val.String())
	}
	return strs
}

func ToIntList(array MagicArray) []int {
	if arr, ok := array.(internal.TArray[int]); ok {
		return []int(arr)
	}
	ints := make([]int, 0, array.Len())
	iter := array.Iter()
	for val := iter.FirstVal(); val != nil; val = iter.NextVal() {
		if intval, ok := val.Int(); ok {
			ints = append(ints, intval)
		}
	}
	return ints
}

func ToAnyList(array MagicArray) []any {
	var ret = make([]any, array.Len())
	iter := array.Iter()
	var i = -1
	for v := iter.FirstVal(); v != nil; v = iter.NextVal() {
		i++
		ret[i] = v.Interface()
	}
	return ret
}

func ToMap(array MagicArray) map[string]any {
	mm := make(map[string]any, array.Len())
	iter := array.Iter()
	for k, v := iter.FirstKV(); k != nil; k, v = iter.NextKV() {
		mm[k.String()] = v.Interface()
	}
	return mm

}
