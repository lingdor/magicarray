package magicarray

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSetColumnTag(t *testing.T) {
	//arr MagicArray,k,v any
	var dbinfo = []struct {
		Id       int
		UserId   int    `json:"userId"`
		UserName string `json:"uname"`
	}{
		{
			Id:       1,
			UserId:   101,
			UserName: "bobby",
		},
		{
			Id:       2,
			UserId:   102,
			UserName: "tom",
		},
	}

	rows, _ := Valueof(dbinfo)
	rows = SetColumnTag(rows, "UserName", "json", "userName")
	if jsonbs, err := json.Marshal(rows); err == nil {
		if string(jsonbs) != `[{"Id":1,"userId":101,"userName":"bobby"},{"Id":2,"userId":102,"userName":"tom"}]` {
			fmt.Println(string(jsonbs))
			t.Fail()
		}
	} else {
		t.Error(err)
	}
}
