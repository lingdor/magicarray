package array

import (
	"bytes"
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/internal"
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
	"io"
	"strings"
	"unicode"
)

func JsonEncode(arr api.IMagicArray, writer io.Writer, opts ...api.JsonOpt) (err error) {
	return internal.JsonEncode(arr, writer, opts...)
}

func JsonMarshal(arr api.IMagicArray, opts ...api.JsonOpt) ([]byte, error) {
	return internal.JsonMarshal(arr, opts...)
}

func JosnOptOmitEmpty(isEmitEmtpty bool) api.JsonOpt {
	return func(info *api.JsonOptInfo) {
		info.ValueFilters = append(info.ValueFilters, func(k api.IZVal, v api.IZVal) (api.IZVal, api.IZVal, bool) {
			if isEmitEmtpty && v.IsNil() {
				return k, v, false
			}
			return k, v, true
		})
	}
}
func JsonOptDefaultNamingUnderscoreToHump() api.JsonOpt {
	return func(info *api.JsonOptInfo) {
		info.NameFilters = append(info.NameFilters,
			func(k api.IZVal, v api.IZVal) (api.IZVal, api.IZVal, bool) {
				if tag, ok := v.(api.ZValTag); ok {
					if _, ok = tag.GetTag("json"); ok {
						//ignore items of json tag have setted
						return k, v, true
					}
				}

				sp := strings.Split(k.String(), "_")
				buf := bytes.Buffer{}
				for i, item := range sp {
					if len(item) < 1 {
						continue
					}
					runes := []rune(item)
					if i == 0 {
						runes[0] = unicode.ToLower(runes[0])
					} else {
						runes[0] = unicode.ToUpper(runes[0])
					}
					buf.Write([]byte(string(runes)))
				}
				return zval.NewZValOfKind(kind.String, buf.String()), v, true
			})
	}
}
