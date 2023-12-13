package api

import "time"

type listType interface {
	~[]string | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | time.Time
}
