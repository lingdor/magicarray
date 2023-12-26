package magicarray

// SetColumnTag Set tags of key column
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

// SetTag Set tag key and value to the value of MagicArray
func SetTag(arr MagicArray, key any, tagk, tagv string) MagicArray {
	return Set(arr, tagk, ZValTagSet(arr.Get(key), tagk, tagv))
}
