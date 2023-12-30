package api

import "time"

type JsonOptInfo struct {
	NameFilters       []JsonOptFilter
	ValueFilters      []JsonOptFilter
	DefaultTimeFormat string
}

func (j *JsonOptInfo) TimeFormat() string {
	if j.DefaultTimeFormat == "" {
		return time.RFC3339
	}
	return j.DefaultTimeFormat
}

type JsonOpt func(*JsonOptInfo)

type JsonOptFilter func(IZVal, IZVal) (IZVal, IZVal, bool)
