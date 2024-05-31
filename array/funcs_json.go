package array

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
	"io"
	"strconv"
	"strings"
	"unicode"
)

func JsonMarshal(arr api.IMagicArray, opts ...api.JsonOpt) ([]byte, error) {

	buff := &bytes.Buffer{}
	optInfo := &api.JsonOptInfo{}
	for _, opt := range opts {
		opt(optInfo)
	}
	err := jsonEncode(arr, buff, optInfo)
	return buff.Bytes(), err

}
func MustJsonMarshal(arr api.IMagicArray, opts ...api.JsonOpt) []byte {
	bs, err := JsonMarshal(arr, opts...)
	if err != nil {
		panic(err)
	}
	return bs
}
func JsonEncode(arr api.IMagicArray, writer io.Writer, opts ...api.JsonOpt) (err error) {
	var bs []byte
	if bs, err = JsonMarshal(arr, opts...); err == nil {
		_, err = writer.Write(bs)
	}
	return
}
func genWrapSpace(size int) []byte {
	bs := make([]byte, size+1)
	bs[0] = '\n'
	for i := 1; i < size; i++ {
		bs[i] = ' '
	}
	return bs
}
func jsonEncode(arr api.IMagicArray, writer io.Writer, optInfo *api.JsonOptInfo) (err error) {

	token := '['
	if arr.IsKeys() {
		token = '{'
	}
	if _, err = writer.Write([]byte{byte(token)}); err != nil {
		return
	}

	iter := arr.Iter()
kloop:
	for k, v := iter.FirstKV(); k != nil; k, v = iter.NextKV() {
		tagName := k.String()
		if tagval, ok := v.(api.ZValTag); ok {
			if readTagVal, ok := tagval.GetTag("json"); ok {
				tagIterator := newJsonTagIterator(readTagVal)
				for tagOptKV := tagIterator.next(); tagOptKV != nil; tagOptKV = tagIterator.next() {
					tagOptK := strings.ToLower(tagOptKV.name)
					if v.IsNil() && tagOptK == "omitempty" {
						continue kloop
					} else if v.Kind() == kind.Time {
						format := optInfo.DefaultTimeFormat
						if tagOptKV.val != "" && tagOptK == "format" {
							format = tagOptKV.val
						}
						if t, ok := v.Time(); ok {
							v = zval.NewZValOfKind(kind.String, t.Format(format))
						} else {
							return fmt.Errorf("convert to time faild, field:%s, value:%s", k.String(), v.String())
						}
					}
				}
				if tagIterator.name != "" {
					tagName = tagIterator.name
				}
			}
		}
		ztag := zval.NewZValOfKind(kind.String, tagName)
		if arr.IsKeys() {
			for _, opt := range optInfo.NameFilters {
				var ok bool
				if ztag, v, ok = opt(ztag, v); !ok {
					continue kloop
				}
			}
		}
		for _, opt := range optInfo.ValueFilters {
			var ok bool
			if ztag, v, ok = opt(ztag, v); !ok {
				continue kloop
			}
		}
		//code

		if optInfo.IndentSize > 0 {
			writer.Write(genWrapSpace((optInfo.Deep + 1) * optInfo.IndentSize))
		}

		var err error
		if arr.IsKeys() {

			if _, err = writer.Write([]byte{byte('"')}); err == nil {
				if _, err = writer.Write([]byte(ztag.String())); err == nil {
					if _, err = writer.Write([]byte("\":")); err == nil {
						err = encodeZval(v, writer, optInfo)
					}
				}
			}

		} else {
			childOpt := *optInfo
			childOpt.Deep++
			err = encodeZval(v, writer, &childOpt)
		}
		if err == nil && iter.Index()+1 < arr.Len() {
			if _, err = writer.Write([]byte(",")); err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}

	if optInfo.IndentSize > 0 {
		if optInfo.Deep > 0 {
			writer.Write(genWrapSpace((optInfo.Deep) * optInfo.IndentSize))
		} else {
			writer.Write([]byte{'\n'})
		}
	}

	token = ']'
	if arr.IsKeys() {
		token = '}'
	}
	if _, err = writer.Write([]byte{byte(token)}); err != nil {
		return
	}
	return err
}

func encodeZval(val api.IZVal, writer io.Writer, optInfo *api.JsonOptInfo) (err error) {
	if val.IsNil() {
		_, err = writer.Write([]byte("null"))
		return
	}
	switch val.Kind() {
	case kind.String, kind.Bytes:
		v := val.String()
		v = strings.ReplaceAll(v, "\"", "\\\"")
		v = strings.ReplaceAll(v, "\n", "\\n")
		if _, err = writer.Write([]byte{byte('"')}); err == nil {
			if _, err = writer.Write([]byte(v)); err == nil {
				_, err = writer.Write([]byte{byte('"')})
			}
		}
		return
	case kind.Bool:
		if v, ok := val.Bool(); ok {
			_, err = writer.Write([]byte(strconv.FormatBool(v)))
			return
		}
	case kind.Int:
		if v, ok := val.Int(); ok {
			_, err = writer.Write([]byte(strconv.Itoa(v)))
			return
		}
	case kind.Uint:
		if v, ok := val.Uint(); ok {
			_, err = writer.Write([]byte(strconv.FormatUint(uint64(v), 10)))
			return
		}
	case kind.Int64:
		if v, ok := val.Int64(); ok {
			_, err = writer.Write([]byte(strconv.FormatInt(v, 10)))
			return
		}
	case kind.Uint64:
		if v, ok := val.Uint64(); ok {
			_, err = writer.Write([]byte(strconv.FormatUint(v, 10)))
			return
		}
	case kind.Float32:
		if v, ok := val.Float32(); ok {
			_, err = writer.Write([]byte(strconv.FormatFloat(float64(v), 'f', -1, 32)))
			return
		}
	case kind.Float64:
		if v, ok := val.Float64(); ok {
			_, err = writer.Write([]byte(strconv.FormatFloat(v, 'f', -1, 64)))
			return
		}
	case kind.MagicArray, kind.Map, kind.Slice, kind.Interface, kind.Struct:
		if arr, ok := val.Arr(); ok {
			return jsonEncode(arr, writer, optInfo)
		}
	}
	_, err = fmt.Fprint(writer, val.Interface())
	if err != nil {
		return
	}
	return
}

type jsonTagOpt struct {
	name string
	val  string
}

type jsonTagIterator struct {
	name   string
	optRaw string
	index  int
}

func (j *jsonTagIterator) next() (kv *jsonTagOpt) {
	start := j.index
	last := start
	var matchQuote byte = 0
	var quoteStart int
	var step uint8 = 0
	jsonOpt := &jsonTagOpt{}

	for ; j.index < len(j.optRaw); j.index++ {
		if matchQuote == 0 && j.optRaw[j.index] == '\'' || j.optRaw[j.index] == '"' {
			matchQuote = j.optRaw[j.index]
			quoteStart = j.index + 1
			continue
		} else if matchQuote != 0 && matchQuote == j.optRaw[j.index] {
			//end
			if step == 0 {
				if quoteStart < len(j.optRaw) {
					jsonOpt.name = j.optRaw[quoteStart : j.index-1]
					continue
				}
			} else if step == 1 {
				if quoteStart < len(j.optRaw) {
					jsonOpt.val = j.optRaw[quoteStart : j.index-1]
					return jsonOpt
				}
			}
		}
		if j.optRaw[j.index] == ',' {
			break
		} else if j.optRaw[j.index] == '=' && step == 0 {
			step = 1
			start = j.index + 1
			jsonOpt.name = j.optRaw[start:j.index]
			continue
		} else {
			isSpace := unicode.IsSpace(rune(j.optRaw[j.index]))
			if j.index == start && isSpace {
				start = j.index + 1
				last = start + 1
				continue
			} else if !isSpace {
				last = j.index + 1
			}
		}
	}
	if step == 1 && start < last && last <= len(j.optRaw) {
		jsonOpt.val = j.optRaw[start:last]
	} else if step == 0 && start < last && last <= len(j.optRaw) {
		jsonOpt.name = j.optRaw[start:last]
	}
	if jsonOpt.name == "" {
		return nil
	}
	return jsonOpt
}

func newJsonTagIterator(optRaw string) *jsonTagIterator {
	iterator := &jsonTagIterator{
		index:  0,
		optRaw: optRaw,
	}
	if next := iterator.next(); next != nil {
		iterator.name = next.name
	}
	return iterator
}

func JsonUnMarshal(content []byte) (MagicArray, error) {
	var val any
	if err := json.Unmarshal(content, &val); err != nil {
		return nil, err
	}
	return Valueof(val)
}
