package array

import (
	"github.com/lingdor/magicarray/internal"
	"github.com/lingdor/magicarray/zval"
)

func init() {
	zval.ToMagicArr = Valueof
	internal.JsonMarshal = JsonMarshal
}
