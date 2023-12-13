package api

type ZValTag interface {
	GetTag(string) (string, bool)
}
type ZValTagSet interface {
	SetTag(s string, val string)
}
