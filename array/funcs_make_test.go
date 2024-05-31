package array

import (
	"testing"
)

func TestStructToMap(t *testing.T) {

	structVal := struct {
		Field1 string
		Field2 int
		Field3 string
	}{
		Field1: "hello",
		Field2: 2,
	}

	expert := map[string]any{
		"Field1": "hello",
		"Field2": 2,
		"Field3": "",
	}

	if arrVal, err := Valueof(structVal); err != nil {
		t.Error(err)
	} else if err := Equals(arrVal, expert); err != nil {
		t.Error(err)
	}
}

func TestSlice(t *testing.T) {

	var xx = []int{123, 1, 2, 3}
	mapArr, err := Valueof(xx)
	if err != nil {
		t.Error(err)
	}
	if err := Equals(mapArr, xx); err != nil {
		t.Error(err)
	}

	strs, _ := Valueof([]string{"123", "456", "780"})

	ints, _ := Valueof([]int{123, 456, 780})

	if err := Equals(strs, ints); err != nil {
		t.Error(err)
	}

}

// //intersect of keys
// keys := []string{"field1", "field2"}
// arr := ValueOf(structVal, NamingJsonFirst)
// arr = IntersectKey(arr, keys)
//
//	if !arr.equals(map[string]interface{}{
//		"field1": "hello",
//		"field2": 2,
//	}) {
//
//		t.Error("intersect_key assert faild")
//	}
//
// //sortMap
// num := 0
//
//	for K, V := range ValueOf(structVal, NamingJsonFirst).sortKeys().reverse().All() {
//		if num == 0 && K == "field3", (V.(string)) == "hello" {
//		} else if num == 1 && K == "field2", V.(int) == 2 {
//
//		} else {
//			t.Error("sortMap faild")
//		}
//		num++
//	}
//
// //array shift test
// arr := ValueOf(structVal, NamingJsonFirst)
// arr = Shift(arr, []string{"field1"})
//
//	if !arr.Equals(map[string]interface{}{
//		"field2": 2,
//		"field3": 3,
//	}) {
//
//		t.Error("array shift test faild")
//	}
//
// //array column
//
//	rows := []map[string]interface{}{
//		{
//			"field1": "a1",
//			"field2": "a2",
//		},
//		{
//			"field1": "b1",
//			"field2": "b2",
//		},
//	}
//
// arr := ValueOf(rows, NamingDefault)
// colArr := arr.Column(arr, "field2")
//
//	if !colArr.Equals([]string{"a2", "b2"}) {
//		t.Error("array column assert faild.")
//	}
//
// nums := []interface{}{"1", 2, 3, "4", 5}
// arr := ValueOf(rows)
//
//	if 15 != arr.Sum(nums) {
//		t.Error("sum faild")
//	}
func TestClone(t *testing.T) {

	arr := ValueofMap(map[string]any{
		"Field1": "hello",
		"Field2": 2,
		"Field3": "",
	})
	newArr := Clone(arr)

	if newArr.Len() != 3 {
		t.Error("cloned array length not expect")
	}
	iter := arr.Iter()
	for k, v := iter.FirstKV(); v != nil; k, v = iter.NextKV() {
		if !newArr.Get(k).Compare(v) {
			t.Errorf(" clone key:%s not equalls", k.String())
		}
	}

}
