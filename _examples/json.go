package main

import (
	"fmt"
	"github.com/lingdor/magicarray/array"
)

func jsonCommand() {
	users, _ := array.Valueof([]map[string]any{
		{
			"id":        1,
			"user_name": "bobby",
			"Age":       nil,
		},
		{
			"id":        2,
			"user_name": "lily",
			"Age":       16,
		},
	})
	bs, _ := array.JsonMarshal(
		users,
		array.JosnOptOmitEmpty(true),
		array.JsonOptDefaultNamingUnderscoreToHump())

	fmt.Println(string(bs))

}
