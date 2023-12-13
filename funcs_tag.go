package magicarray

import "unicode"

func SetColumnTag(array MagicArray, columnKey any, tagk, tagv string) MagicArray {

	iter := array.Iter()
	for rowk, row := iter.FirstKV(); row != nil; rowk, row = iter.NextKV() {
		if rowArr, ok := row.Arr(); ok {
			newArr := Set(rowArr, columnKey, ZValTagSet(rowArr.Get(columnKey), tagk, tagv))
			if newArr != array {
				array = Set(array, rowk, rowArr)
			}
		}
	}
	return array
}

func WashColumnTag(array MagicArray, rules ...WashTagRule) MagicArray {

	iter := array.Iter()
	for rowk, row := iter.FirstKV(); row != nil; rowk, row = iter.NextKV() {
		if rowArr, ok := row.Arr(); ok {
			newArr := WashTag(rowArr, rules...)
			if newArr != rowArr {
				array = Set(array, rowk, newArr)
			}
		}
	}
	return array
}

func SetTag(arr MagicArray, key any, tagk, tagv string) MagicArray {
	return Set(arr, tagk, ZValTagSet(arr.Get(key), tagk, tagv))
}

type WashTagRule func(key, val ZVal) ZVal

func WashTag(arr MagicArray, rules ...WashTagRule) MagicArray {

	var writer MagicArray = ToWriter(arr)
	iter := writer.Iter()
	for k, v := iter.FirstKV(); k != nil; k, v = iter.NextKV() {
		for _, rule := range rules {
			if newzval := rule(k, v); newzval != v {
				writer = Set(writer, k, newzval)
			}
		}
	}
	return writer
}

func WashTagJsonInitalLowerOpt() WashTagRule {

	return func(k, v ZVal) ZVal {
		if _, ok := ZValTagGet(v, "json"); !ok {
			runes := []rune(k.String())
			if len(runes) < 1 {
				return v
			}
			runes[0] = unicode.ToLower(runes[0])
			return ZValTagSet(v, "json", string(runes))
		}
		return v
	}

}
