package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lingdor/magicarray/api"
	"github.com/lingdor/magicarray/kind"
	"github.com/lingdor/magicarray/zval"
	"io"
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
func JsonEncode(arr api.IMagicArray, writer io.Writer, opts ...api.JsonOpt) (err error) {
	var bs []byte
	if bs, err = JsonMarshal(arr, opts...); err == nil {
		_, err = writer.Write(bs)
	}
	return
}
func jsonEncode(arr api.IMagicArray, writer io.Writer, optInfo *api.JsonOptInfo) (err error) {

	var first = true
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
		if !first {
			if _, err := writer.Write([]byte(",")); err != nil {
				return err
			}
		} else {
			first = false
		}
		var err error
		if arr.IsKeys() {
			if _, err = writer.Write([]byte(fmt.Sprintf("\"%s\":", ztag.String()))); err == nil {
				err = encodeZval(v, writer, optInfo)
			}
		} else {
			err = encodeZval(v, writer, optInfo)
		}
		if err != nil {
			return err
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
	if arr, ok := val.Arr(); ok {
		return jsonEncode(arr, writer, optInfo)
	}
	var bs []byte
	if bs, err = json.Marshal(val.Interface()); err == nil {
		_, err = writer.Write(bs)
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
				last = start
				continue
			} else if !isSpace {
				last = j.index
			}
		}
	}
	if step == 1 && start < len(j.optRaw) && last < len(j.optRaw) {
		jsonOpt.val = j.optRaw[start : last+1]
	}
	return jsonOpt
}

func newJsonTagIterator(optRaw string) *jsonTagIterator {
	iterator := &jsonTagIterator{
		index:  -1,
		optRaw: optRaw,
	}
	if next := iterator.next(); next != nil {
		iterator.name = next.name
	}
	return iterator
}
