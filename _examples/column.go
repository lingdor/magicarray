package main

import (
	"encoding/json"
	"fmt"
	"github.com/lingdor/magicarray/array"
)

type ColumnUserEntity struct {
	Id       int `json:"uid"`
	UserName string
	IsMale   bool
}

func columnCommand() {

	users := []ColumnUserEntity{
		{
			Id:       1,
			UserName: "Bobby",
			IsMale:   true,
		},
		{
			Id:       2,
			UserName: "Lily",
			IsMale:   false,
		},
	}

	usersArr := array.ValueOfSlice(users)
	usersArr = array.WashAll(usersArr, array.WashTagRuleJsonInitialLower())
	if bs, err := json.Marshal(usersArr); err == nil {
		fmt.Println(string(bs))
	} else {
		panic(err)
	}

	usersArr = array.Column(usersArr, "UserName")
	if bs, err := json.Marshal(usersArr); err == nil {
		fmt.Println(string(bs))
	} else {
		panic(err)
	}

}
