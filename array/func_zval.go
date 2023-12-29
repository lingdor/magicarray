package array

import (
	"github.com/lingdor/magicarray/api"
	zval2 "github.com/lingdor/magicarray/zval"
)

func ZValTagSet(z ZVal, tag, val string) ZVal {

	if setter, ok := z.(api.ZValTagSet); ok {
		setter.SetTag(tag, val)
		return z
	}
	z = zval2.NewTagZVal(z)
	setter, _ := z.(api.ZValTagSet)
	setter.SetTag(tag, val)
	return z
}
func ZValTagGet(z ZVal, tag string) (string, bool) {

	if ztag, ok := z.(api.ZValTag); ok {
		return ztag.GetTag(tag)
	}
	return "", false
}
