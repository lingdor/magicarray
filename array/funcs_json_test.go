package array

import (
	"testing"
)

func TestJsonUnMarhsal(t *testing.T) {

	var json = `
			[
			{"subject":"new1","content":{"num":123,"bl":true,"str":"abc"}},
			{"subject":"new1","content":{"num":456,"bl":false,"str":"def"}}
			]
		`

	arr, _ := JsonUnMarshal([]byte(json))

	if arr.Get(0).MustArr().Get("content").MustArr().Get("num").MustInt() != 123 {
		t.Errorf("JsonUnMarshal not assert [0].content.num == 123")
	}

	json = `null`
	arr, _ = JsonUnMarshal([]byte(json))
	if arr != nil {
		t.Errorf("JsonUnMarshal is not assert")
	}

}
