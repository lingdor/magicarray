package array

import (
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
	"unicode"
)

// WashColumn Wash the value of MagicArray column by rules
func WashColumn(array MagicArray, column string, rules ...WashRuleFunc) MagicArray {

	newArr := Make(false, false, array.Len())
	iter := array.Iter()
rowLoop:
	for row := iter.FirstVal(); row != nil; row = iter.NextVal() {
		if rowArr, ok := row.Arr(); ok {
			rowArr = ToWriter(rowArr)
			for _, rule := range rules {
				oldval := rowArr.Get(column)
				if newk, newv, ok := rule(zval.NewZValOfKind(kind.String, column), oldval); ok {
					if newk.Compare(zval.NewZValOfKind(kind.String, column)) {
						if !newv.Compare(oldval) {
							rowArr = Set(rowArr, newk, newv)
						}
					} else {
						rowArr = Remove(rowArr, column)
						if newk != nil && !newk.IsNil() {
							rowArr = Set(rowArr, newk, newv)
						}
					}
				} else {
					continue rowLoop
				}
				newArr = Append(newArr, rowArr)
			}
		}
	}
	return array
}

// WashRuleFunc type  of wash rule function
type WashRuleFunc func(key, val ZVal) (ZVal, ZVal, bool)

// WashAll Wash the value of MagicArray all values by rules
func WashAll(arr MagicArray, rules ...WashRuleFunc) MagicArray {

	newArr := Make(true, true, arr.Len())
	iter := arr.Iter()
rowLoop:
	for k, v := iter.FirstKV(); k != nil; k, v = iter.NextKV() {
		for _, rule := range rules {
			if newk, newv, ok := rule(k, v); ok {
				newArr = Set(newArr, newk, newv)
			} else {
				continue rowLoop
			}
		}
	}
	return newArr
}

// WashTagRuleJsonInitialLower Wash the value tags ,lower the initial letter if no fund the json tag.
func WashTagRuleJsonInitialLower() WashRuleFunc {

	return func(k, v ZVal) (ZVal, ZVal, bool) {
		if _, ok := ZValTagGet(v, "json"); !ok {
			runes := []rune(k.String())
			if len(runes) < 1 {
				return k, v, true
			}
			runes[0] = unicode.ToLower(runes[0])
			return k, ZValTagSet(v, "json", string(runes)), true
		}
		return k, v, true
	}

}
