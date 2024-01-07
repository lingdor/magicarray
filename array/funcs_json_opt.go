package array

import (
	"github.com/lingdor/magicarray/api"
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
func JsonOptDefaultNamingUnderscoreToCamelCase() api.JsonOpt {
	return func(info *api.JsonOptInfo) {
		washopt := GetWashFuncWashUnderScoreCaseToCamelCase(false)
		info.NameFilters = append(info.NameFilters,
			func(k api.IZVal, v api.IZVal) (api.IZVal, api.IZVal, bool) {
				if tag, ok := v.(api.ZValTag); ok {
					if _, ok = tag.GetTag("json"); ok {
						//ignore items of json tag have setted
						return k, v, true
					}
				}
				return washopt(k, v)
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
