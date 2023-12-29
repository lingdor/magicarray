package api

type MagicArray interface {
	Iter() Iterator
	KV
	Len
	Getter
}

type Iterator interface {
	// All The iterate of Array
	NextKV() (ZVal, ZVal)
	FirstKV() (ZVal, ZVal)
	//NextKey() ZVal todo
	NextVal() ZVal
	FirstVal() ZVal
	Index() int
}
type Len interface {
	Len() int
}

type KV interface {
	Keys() MagicArray
	Values() MagicArray
	IsKeys() bool
}
type Getter interface {
	Get(key interface{}) ZVal
}
type Setter interface {
	Set(key interface{}, val interface{}) WriteMagicArray
}
type Appender interface {
	Append(val any) WriteMagicArray
}
type WriteMagicArray interface {
	MagicArray
	Setter
	Appender
	Remover
}

type Remover interface {
	Remove(key any) (WriteMagicArray, error)
}
