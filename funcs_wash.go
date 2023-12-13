package magicarray

func Wash(array MagicArray, opts ...Opt) MagicArray {
	var ret MagicArray = array
	for _, opt := range opts {
		ret = opt(array)
	}
	return ret
}
