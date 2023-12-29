package array

import (
	"encoding/json"
	"fmt"
	"github.com/lingdor/magicarray/api"
	"testing"
)

func TestJson(t *testing.T) {

	type tt struct {
		UserId   int
		UserName *string
		Ext      any `json:"ext,omitempty"`
	}

	users := []tt{
		{UserId: 1, UserName: nil, Ext: 0},
		{UserId: 1, UserName: nil, Ext: ""},
		{UserId: 1, UserName: nil, Ext: 0.0},
		{UserId: 1, UserName: nil, Ext: nil},
	}
	var err error
	var bs2 []byte
	var arr api.IMagicArray

	//if bs, err = json.Marshal(users); err == nil {
	if arr, err = Valueof(users); err == nil {
		arr = SetColumnTag(arr, "Ext", "ff", "11")
		if bs2, err = json.Marshal(arr); err == nil {
			fmt.Println(string(bs2))
			return
		}
	}
	//}
	t.Error(err)
}
