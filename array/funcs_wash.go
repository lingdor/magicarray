package array

import "unicode"

// WashColumn Wash the value of MagicArray column by rules
func WashColumn(array MagicArray, rules ...WashRuleFunc) MagicArray {

	iter := array.Iter()
	for rowk, row := iter.FirstKV(); row != nil; rowk, row = iter.NextKV() {
		if rowArr, ok := row.Arr(); ok {
			newArr := WashAll(rowArr, rules...)
			if newArr != rowArr {
				array = Set(array, rowk, newArr)
			}
		}
	}
	return array
}

// WashRuleFunc type  of wash rule function
type WashRuleFunc func(key, val ZVal) (ZVal, bool)

// WashAll Wash the value of MagicArray all values by rules
func WashAll(arr MagicArray, rules ...WashRuleFunc) MagicArray {

	var writer MagicArray = ToWriter(arr)
	iter := writer.Iter()
	for k, v := iter.FirstKV(); k != nil; k, v = iter.NextKV() {
		for _, rule := range rules {
			if newzval, ok := rule(k, v); ok && newzval != v {
				writer = Set(writer, k, newzval)
			} else if !ok {
				writer = Remove(writer, k)
			}
		}
	}
	return writer
}

// WashTagRuleJsonInitialLower Wash the value tags ,lower the initial letter if no fund the json tag.
func WashTagRuleJsonInitialLower() WashRuleFunc {

	return func(k, v ZVal) (ZVal, bool) {
		if _, ok := ZValTagGet(v, "json"); !ok {
			runes := []rune(k.String())
			if len(runes) < 1 {
				return v, true
			}
			runes[0] = unicode.ToLower(runes[0])
			return ZValTagSet(v, "json", string(runes)), true
		}
		return v, true
	}

}
