package magicarray

import (
	"errors"
	"fmt"
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
)

func Equals(from MagicArray, to any) error {
	toArr, err := Valueof(to)
	if err != nil {
		return err
	}
	if from.Len() != toArr.Len() {
		return errors.New(fmt.Sprintf("magicarray length is not equals, from=%d,to=%d", from.Len(), toArr.Len()))
	}
	fromIter := from.Iter()
	for k, v := fromIter.FirstKV(); k != nil && v != nil; k, v = fromIter.NextKV() {
		toV := toArr.Get(k)
		if !v.Compare(toV) {
			return errors.New(fmt.Sprintf("magicarray.value is not equals, key=%s,value:'%s'->'%s'", k.String(), v.String(), toV.String()))
		}
	}
	return nil
}
func Max(arr MagicArray) ZVal {
	maxVal := zval.NewZValNil()
	iter := arr.Iter()
	for val := iter.FirstVal(); val != nil; val = iter.NextVal() {
		if val.Kind() == kind.Int && maxVal.Kind() == kind.Int && val.MustInt() > maxVal.MustInt() {
			maxVal = val
		} else if val.Kind() == kind.Int64 && maxVal.Kind() == kind.Int64 && val.MustInt64() > maxVal.MustInt64() {
			maxVal = val
		}
		if ff, ok := val.Float64(); ok {
			if maxFloat, ok := maxVal.Float64(); ok {
				if maxVal.IsNil() || maxFloat < ff {
					maxVal = val
				}
			}
		}
	}
	return maxVal
}
func Min(arr MagicArray) ZVal {

	minVal := zval.NewZValNil()
	iter := arr.Iter()
	for val := iter.FirstVal(); val != nil; val = iter.NextVal() {
		if val.Kind() == kind.Int && minVal.Kind() == kind.Int && val.MustInt() < minVal.MustInt() {
			minVal = val
		} else if val.Kind() == kind.Int64 && minVal.Kind() == kind.Int64 && val.MustInt64() < minVal.MustInt64() {
			minVal = val
		}
		if ff, ok := val.Float64(); ok {
			if maxFloat, ok := minVal.Float64(); ok {
				if minVal.IsNil() || maxFloat > ff {
					minVal = val
				}
			}
		}
	}
	return minVal
}
func Sum(arr MagicArray) ZVal {

	retVal := zval.NewZValOfKind(kind.Int, 0)
	iter := arr.Iter()
	for val := iter.FirstVal(); val != nil; val = iter.NextVal() {
		if val.Kind() == kind.Int && retVal.Kind() == kind.Int {
			retVal = zval.NewZValOfKind(kind.Int, retVal.MustInt()+val.MustInt())
		} else if val.Kind() == kind.Int64 && retVal.Kind() == kind.Int64 && val.MustInt64() < retVal.MustInt64() {
			retVal = zval.NewZValOfKind(kind.Int64, retVal.MustInt()+val.MustInt())
		}
		if valFloat, ok := val.Float64(); ok {
			if sumFloat, ok := retVal.Float64(); ok {
				retVal = zval.NewZValOfKind(kind.Float64, sumFloat+valFloat)
			}
		}
	}
	return retVal
}

// In check value is in MagicArray
func In(arr MagicArray, value any) bool {

	iter := arr.Iter()
	for val := iter.FirstVal(); val != nil; val = iter.NextVal() {
		if val.Compare(zval.NewZVal(value)) {
			return true
		}
	}
	return false
}
