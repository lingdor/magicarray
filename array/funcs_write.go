package array

// Merge fields from parameters to MagicArray
func Merge(arr MagicArray, args ...any) (MagicArray, error) {

	var setter MagicArray = ToWriter(arr)
	for _, arg := range args {
		toArr, err := Valueof(arg)
		if err != nil {
			return nil, err
		}
		iter := toArr.Iter()
		if toArr.IsKeys() {
			for k, v := iter.FirstKV(); k != nil; k, v = iter.NextKV() {
				setter = Set(setter, k, v)
			}
		} else {
			for v := iter.FirstVal(); v != nil; v = iter.NextVal() {
				setter = Append(setter, v)
			}
		}
	}
	return setter, nil
}

// Append value to Magic
func Append(arr MagicArray, val any) MagicArray {
	return ToWriter(arr).Append(val)
}

// Set value of MagicArray
func Set(arr MagicArray, key, val any) MagicArray {
	setter := ToWriter(arr)
	setter = setter.Set(key, val)
	return setter
}

// Remove item from the MagicArray
func Remove(arr MagicArray, keys ...any) MagicArray {

	writeArr := ToWriter(arr)
	for _, key := range keys {
		writeArr, _ = writeArr.Remove(key)
	}
	return writeArr
}
