package array

import (
	"bytes"
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
	"strings"
	"unicode"
)

func GetWashKeysFuncWashUnderScoreCaseToCamelCase(firstUpper bool) WashRuleFunc {
	return func(k api.IZVal, v api.IZVal) (api.IZVal, api.IZVal, bool) {
		sp := strings.Split(strings.ToLower(k.String()), "_")
		buf := bytes.Buffer{}
		for i, item := range sp {
			if len(item) < 1 {
				continue
			}
			runes := []rune(item)
			if firstUpper || i != 0 {
				runes[0] = unicode.ToUpper(runes[0])
			}
			buf.Write([]byte(string(runes)))
		}
		return zval.NewZValOfKind(kind.String, buf.String()), v, true
	}
}

func GetWashFuncWashUnderScoreCaseToCamelCase(firstUpper bool) WashRuleFunc {
	return func(k api.IZVal, v api.IZVal) (api.IZVal, api.IZVal, bool) {
		sp := strings.Split(strings.ToLower(v.String()), "_")
		buf := bytes.Buffer{}
		for i, item := range sp {
			if len(item) < 1 {
				continue
			}
			runes := []rune(item)
			if firstUpper || i != 0 {
				runes[0] = unicode.ToUpper(runes[0])
			}
			buf.Write([]byte(string(runes)))
		}
		return k, zval.NewZValOfKind(kind.String, buf.String()), true
	}
}
