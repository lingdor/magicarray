package api

type TagVal interface {
	Name() string
	Iterator
	SetOpt(string, string)
}
