package array

import (
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/zval"
)

func init() {
	zval.ToMagicArr = func(list any) (api.MagicArray, error) {
		return Valueof(list)
	}
}
