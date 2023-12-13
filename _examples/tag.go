package main

import (
	"encoding/json"
	"fmt"
	arr "github.com/lingdor/magicarray"
)

type UserEntity struct {
	Id       int `json:"uid"`
	UserName string
	IsMale   bool
}

func tagCommand() {

	users := []UserEntity{
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

	usersArr := arr.ValueOfSlice(users)
	usersArr = arr.WashColumnTag(usersArr, arr.WashTagJsonInitalLowerOpt())
	if bs, err := json.Marshal(usersArr); err == nil {
		fmt.Println(string(bs))
	} else {
		panic(err)
	}

}
