package array

import "github.com/lingdor/magicarray/api"

// ZVal is the mixed value item of MagicArray, You can get int, string, bool etc. type of value by ZVal.
// for example:
// zval.String()
// zval.IsEmpty()
// zval.IsNil
// zval.Int()
// zval.Arr()
// ...
type ZVal = api.IZVal
