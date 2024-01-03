package array

import (
	"bytes"
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
	"strings"
	"unicode"
)

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

				sp := strings.Split(strings.ToLower(k.String()), "_")
				buf := bytes.Buffer{}
				for i, item := range sp {
					if len(item) < 1 {
						continue
					}
					runes := []rune(item)
					if i != 0 {
						runes[0] = unicode.ToUpper(runes[0])
					}
					buf.Write([]byte(string(runes)))
				}
				return zval.NewZValOfKind(kind.String, buf.String()), v, true
			})
	}
}
func JsonOptIndent(size int) api.JsonOpt {
	return func(info *api.JsonOptInfo) {
		info.IndentSize = size
	}
}
func JsonOptIndent4() api.JsonOpt {
	return func(info *api.JsonOptInfo) {
		info.IndentSize = 4
	}
}
