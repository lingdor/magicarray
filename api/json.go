package api

type JsonOptInfo struct {
	NameFilters  []JsonOptFilter
	ValueFilters []JsonOptFilter
}

type JsonOpt func(*JsonOptInfo)

type JsonOptFilter func(IZVal, IZVal) (IZVal, IZVal, bool)
