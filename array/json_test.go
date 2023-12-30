package array

import (
	"encoding/json"
	"fmt"
	"testing"
)

func genMap(size int) map[string]any {
	var mm = make(map[string]any, size)
	for i := 0; i < size; i++ {
		key := fmt.Sprintf("index%d", i)
		mm[key] = i
		key = fmt.Sprintf("str%d", i)
		mm[key] = key
	}
	return mm
}

func BenchmarkMagicArrayMapJson(b *testing.B) {

	mm := genMap(50)
	if arr, err := Valueof(mm); err == nil {
		for i := 0; i < b.N; i++ {
			if _, err := JsonMarshal(arr); err != nil {
				b.Error(err)
			}
		}
	} else {
		b.Error(err)
	}
}
func BenchmarkMapJson(b *testing.B) {

	mm := genMap(50)
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(mm); err != nil {
			b.Error(err)
		}

	}
}
func TestXX(t *testing.T) {

	mm := genMap(50)
	if arr, err := Valueof(mm); err == nil {
		if _, err := JsonMarshal(arr); err != nil {
			t.Error(err)
		}
	} else {
		t.Error(err)
	}

}
