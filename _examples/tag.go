package main

import (
	"encoding/json"
	"fmt"
	"github.com/lingdor/magicarray/array"
)

type UserEntity struct {
	Id       int `json:"uid"`
	UserName string
	IsMale   bool
}

func tagCommand() {

	users := UserEntity{
		Id:       1,
		UserName: "Bobby",
		IsMale:   true,
	}

	userArr := arr.ValueofStruct(users)
	userArr = arr.SetTag(userArr, "Id", "json", "UserId")
	if bs, err := json.Marshal(userArr); err == nil {
		fmt.Println(string(bs))
	} else {
		panic(err)
	}

}
