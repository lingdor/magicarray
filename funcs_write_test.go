package magicarray

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {

	structVal := struct {
		Field1 string
		Field2 int
		Field3 *string
	}{
		Field1: "hello",
		Field2: 2,
	}

	mapVal := make(map[string]interface{})
	mapVal["hello"] = 123
	mapVal["Field2"] = 16

	expect := map[string]any{
		"Field2": 16,
		"Field1": "hello",
		"hello":  123,
	}

	//merge
	arr, _ := Valueof(structVal)
	arr, _ = Merge(arr, mapVal)
	if err := Equals(arr, expect); err != nil {
		fmt.Println("array:")
		fmt.Println(ToJson(arr))
		fmt.Println("expect:")
		fmt.Println(ToJson(MustValueof(expect)))
		t.Error("merge assert faild", err)
	}

}

func ToJson(array MagicArray) string {
	bs, err := json.Marshal(array)
	if err != nil {
		panic(err)
	}
	return string(bs)
}
