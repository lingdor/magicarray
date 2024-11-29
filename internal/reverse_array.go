package internal

import "github.com/lingdor/magicarray/api"

type ReverseMap struct {
	arr api.IMagicArray
}

func (r *ReverseMap) Iter_() (_ api.Iterator) {
	return r.arr.RIter_()
}

func (r *ReverseMap) RIter_() (_ api.Iterator) {
	return r.arr.Iter_()
}

func (r *ReverseMap) IterRows() (_ api.IterRowsFunc) {
	return func(yield func(api.IZVal, api.IMagicArray) bool) {
		for k, v := range r.Iter() {
			if !yield(k, v.MustArr()) {
				return
			}
		}
	}
}

func (r *ReverseMap) Iter() (_ api.IterFunc) {
	return func(yield func(api.IZVal, api.IZVal) bool) {
		iter_ := r.arr.RIter_()
		for k, v := iter_.FirstKV(); k != nil; k, v = iter_.NextKV() {
			yield(k, v)
		}
	}
}

func (r *ReverseMap) Keys() (_ api.IMagicArray) {

	panic("todo")

}

func (r *ReverseMap) Values() (_ api.IMagicArray) {

	panic("todo")
}

func (r *ReverseMap) IsKeys() bool {
	return r.IsKeys()
}

func (r *ReverseMap) Len() int {
	return r.arr.Len()
}

func (r *ReverseMap) Get(key interface{}) api.IZVal {

	if r.arr.IsKeys() {
		return r.arr.Get(key)
	}
	idx := key.(int)
	len := r.arr.Len()
	return r.arr.Values().Get(len - idx - 1)
}

func (r *ReverseMap) Reverse() api.IMagicArray {
	return r.arr
}
