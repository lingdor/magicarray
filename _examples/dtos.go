package main

import (
	"encoding/json"
	"fmt"
	"github.com/lingdor/magicarray/array"
	"time"
)

type UserDTO struct {
	Id   int `json:"userid"`
	Name string
}

type ScoreDTO struct {
	Score     int
	ScoreTime time.Time
}

type AreaDto struct {
	CityId int
	City   string
}

func dtosCommand() {

	user := UserDTO{
		Id:   1,
		Name: "bobby",
	}
	score := ScoreDTO{
		Score:     66,
		ScoreTime: time.Now(),
	}
	area := AreaDto{
		CityId: 10000,
		City:   "beij",
	}

	mix, _ := arr.Merge(arr.ValueofStruct(user), score, area)
	mix = arr.Pick(mix, "Id", "City", "Score")
	if bs, err := json.Marshal(mix); err == nil {
		fmt.Println(string(bs))
	} else {
		panic(err)
	}

}
