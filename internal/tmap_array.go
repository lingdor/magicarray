package internal

import (
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/zval"
)

type TMapArray map[string]any

func (t TMapArray) IsKeys() bool {
	return true
}

func (t TMapArray) Keys() api.IMagicArray {
	keys := t.genKeys()
	return TArray[string](keys)
}

func (t TMapArray) Values() api.IMagicArray {
	var vals = make([]any, t.Len())
	i := -1
	for _, v := range t {
		i++
		vals[i] = v
	}
	return TArray[any](vals)
}

func (t TMapArray) Len() int {
	return len(t)
}

func (t TMapArray) Get(key any) api.IZVal {
	var ok bool
	var strKey string
	if strKey, ok = key.(string); ok {
	} else if zvalKey, ok := key.(api.IZVal); ok {
		strKey = zvalKey.String()
	} else {
		strKey = zval.NewZVal(key).String()
	}
	rawVal := t[strKey]
	return zval.NewZVal(rawVal)
}
func (t TMapArray) genKeys() []string {

	keys := make([]string, t.Len())
	i := -1
	for key, _ := range t {
		i++
		keys[i] = key
	}
	return keys
}

func (t TMapArray) Iter() api.Iterator {

	return &TMapArrayIterator{
		arr:   t,
		index: 0,
		keys:  t.genKeys(),
	}
}

func (t TMapArray) MarshalJSON() ([]byte, error) {
	return JsonMarshal(t)
}
