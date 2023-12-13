package errs

import "errors"

var TypeAssertError = errors.New("Magicarray type assert faild")
var LinksDeepOutError = errors.New("links deep out ")
var NoFundKey = errors.New("no found keys")
var OutOfArrayLength = errors.New("index out of array length")
