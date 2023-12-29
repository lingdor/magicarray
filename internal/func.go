package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
	"io"
	"reflect"
	"strings"
)

func GenListKeys(len int) []int {
	var keys = make([]int, len)
	for i := 0; i < len; i++ {
		keys[i] = i
	}
	return keys
}

func SlicetoAnyList(refVal reflect.Value) []any {
	len := refVal.Len()
	ret := make([]any, 0, len)
	for i := 0; i < len; i++ {
		ret = append(ret, refVal.Index(i).Interface())
	}
	return ret
}
func newTArray[T any](listVal []T) api.IMagicArray {
	return TArray[T](listVal)
}

func JsonMarshal(arr api.IMagicArray, opts ...api.JsonOpt) ([]byte, error) {

	buff := &bytes.Buffer{}
	err := JsonEncode(arr, buff, opts...)
	return buff.Bytes(), err

}
func JsonEncode(arr api.IMagicArray, writer io.Writer, opts ...api.JsonOpt) (err error) {
	optInfo := &api.JsonOptInfo{}
	for _, opt := range opts {
		opt(optInfo)
	}
	return jsonEncode(arr, writer, optInfo)
}
func jsonEncode(arr api.IMagicArray, writer io.Writer, optInfo *api.JsonOptInfo) (err error) {

	iter := arr.Iter()
	var first = true
	if arr.IsKeys() {
		writer.Write([]byte("{"))
	kloop:
		for k, v := iter.FirstKV(); k != nil; k, v = iter.NextKV() {

			tagName := k.String()
			if tagval, ok := v.(api.ZValTag); ok {
				if readTagName, ok := tagval.GetTag("json"); ok {
					readTagNames := strings.Split(readTagName, ",")
					for i := 1; i < len(readTagNames); i++ {
						if v.IsNil() && strings.ToLower(strings.TrimSpace(readTagNames[i])) == "omitempty" {
							continue kloop
						}
					}
					tagName = readTagNames[0]
				}
			}
			ztag := zval.NewZValOfKind(kind.String, tagName)
			for _, opt := range optInfo.NameFilters {
				var ok bool
				if ztag, v, ok = opt(ztag, v); !ok {
					continue kloop
				}
			}
			for _, opt := range optInfo.ValueFilters {
				var ok bool
				if ztag, v, ok = opt(ztag, v); !ok {
					continue kloop
				}
			}

			if !first {
				if _, err := writer.Write([]byte(",")); err != nil {
					return err
				}
			} else {
				first = false
			}
			var err error
			if _, err = writer.Write([]byte(fmt.Sprintf("\"%s\":", ztag.String()))); err == nil {
				err = encodeZval(v, writer, optInfo)
			}
			if err != nil {
				return err
			}
		}
		_, err := writer.Write([]byte("}"))
		return err
	}

	if _, err = writer.Write([]byte("[")); err == nil {
	sloop:
		for v := iter.FirstVal(); v != nil; v = iter.NextVal() {

			for _, opt := range optInfo.ValueFilters {
				var ok bool
				if _, v, ok = opt(zval.NewZValOfKind(kind.Int, iter.Index()), v); !ok {
					continue sloop
				}
			}

			if !first {
				if _, err := writer.Write([]byte(",")); err != nil {
					return err
				}
			} else {
				first = false
			}
			if err = encodeZval(v, writer, optInfo); err != nil {
				return
			}
		}
	}
	_, err = writer.Write([]byte("]"))
	return
}
func encodeZval(val api.IZVal, writer io.Writer, optInfo *api.JsonOptInfo) (err error) {
	if arr, ok := val.Arr(); ok {
		return jsonEncode(arr, writer, optInfo)
	}
	var bs []byte
	if bs, err = json.Marshal(val.Interface()); err == nil {
		_, err = writer.Write(bs)
	}
	return
}

//
//func NamingJsonHump(arr api.IMagicArray) api.IMagicArray {
//
//	if arr == nil {
//		return &ZValArray{isKeys: false, listVals: []api.IZVal{}}
//	}
//	if arr.IsKeys() {
//		var newKeys = make([]string, arr.Len())
//		var newVals = make(map[string]ZValArrayMapVal, arr.Len())
//		var refType api.RefType
//		var refTypeOK bool
//		refType, refTypeOK = arr.(api.RefType)
//		var iter = arr.Iter()
//		var i = -1
//		for k, v := iter.FirstKV(); k != nil; k, v = iter.NextKV() {
//			i++
//			ismatch := false
//			if refTypeOK {
//				tt := refType.GetRefType(k.String())
//				if jsonName, ok := tt.Tag.Lookup("json"); ok {
//					newKeys[i] = jsonName
//					ismatch = true
//				}
//			}
//			if !ismatch && len(k.String()) > 0 {
//				runes := []rune(k.String())
//				runes[0] = unicode.ToLower(runes[0])
//				newKeys[i] = string(runes)
//			}
//			newVals[newKeys[i]] = ZValArrayMapVal{val: v, index: i}
//		}
//		return &ZValArray{
//			keys:     newKeys,
//			isSorted: true,
//			isKeys:   true,
//			mapVals:  newVals,
//		}
//	}
//
//	vals := make([]api.IZVal, arr.Len())
//	for i := 0; i < arr.Len(); i++ {
//		v := arr.Get(i)
//		if child, ok := v.Arr(); ok {
//			vals[i] = zval.NewZValOfKind(kind._MagicArray, NamingJsonHump(child))
//		} else {
//			vals[i] = v
//		}
//	}
//	return TArray[api.IZVal](vals)
//
//}
