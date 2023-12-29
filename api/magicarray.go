package api

type IMagicArray interface {
	Iter() Iterator
	KV
	Len
	Getter
}

type Iterator interface {
	// All The iterate of Array
	NextKV() (IZVal, IZVal)
	FirstKV() (IZVal, IZVal)
	//NextKey() IZVal todo
	NextVal() IZVal
	FirstVal() IZVal
	Index() int
}
type Len interface {
	Len() int
}

type KV interface {
	Keys() IMagicArray
	Values() IMagicArray
	IsKeys() bool
}
type Getter interface {
	Get(key interface{}) IZVal
}
type Setter interface {
	Set(key interface{}, val interface{}) WriteMagicArray
}
type Appender interface {
	Append(val any) WriteMagicArray
}
type WriteMagicArray interface {
	IMagicArray
	Setter
	Appender
	Remover
}

type Remover interface {
	Remove(key any) (WriteMagicArray, error)
}
