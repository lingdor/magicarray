# MagicArray 
NO care type, no care struct, easy to sort and aggregation.similar with php array structure,easy to deal for data process. Coding shortly,process automatically.
in magicarray, no nil forever.
```shell
go get github.com/lingdor/magicarray
```
# Functions
| Name                        | Describe                                                                            |
|-----------------------------|-------------------------------------------------------------------------------------|
| Makeof                      | Make a MagicArray instance from struct,slice,array,map                              |
| ValueofStruct               | Make a MagicArray instance from struct, performance better than Makeof.             |
| ValueOfSlice                | Make a MagicArray instance from Slice or array, performance better than Makeof      |
| ValueofStructs              | Make a MagicArray instance from Structs, performance better than Makeof             |
| MustValueof                 | Make a MagicArray instance like Makeof,If return error than panic.                  |
| Make                        | Make a empty MagicArray instance                                                    |
| Equals                      | Return If lenth equal and key and value equals                                      |
| Max                         | Get the maximum  numberic value from the array                                      |
| Min                         | Get the minmum numeric value from the array                                         |
| Sum                         | Calculate sum the total numeric  values from the array                              |
| In                          | check value is in MagicArray                                                        |
| ToStringList                | translate the MagicArray to string array                                            |
| ToIntList                   | translate the MagicArray to integer array                                           |
| ToAnyList                   | translate the MagicArray to any type of array                                       |
| ToMap                       | translate the MagicArray to map                                                     |
| Column                      | Pick the Column from the two-dimensional table data                                 |
| Len                         | Get length of the MagicArray                                                        | 
| Get                         | Get item from the MagicArray                                                        |
| Keys                        | Get keys of the MagicArray                                                          |
| Values                      | Get values of the MagicArray                                                        |
| Pick                        | Pick the keys and values to a new MagicArray for parameter keys order               |
| SetColumnTag                | Set tags of key column                                                              |
| WashColumn                  | Wash the value of MagicArray column by rules                                        |
| SetTag                      | Set tag key and value to the value of MagicArray                                    |
| WashAll                     | Wash the value of MagicArray all values by rules                                    |
| WashTagRuleJsonInitialLower | Wash the value tags ,lower the initial letter if no fund the json tag               |
| Merge                       | Merge fields from parameters to MagicArray                                          |
| Append                      | Append value to Magic                                                               |
| Set                         | Set value of MagicArray                                                             | 
| Remove                      | Remove item from the MagicArray                                                     |
| Implode                     | join MagicArray values to a string                                                  |
| Explode                     | split the string to MagicArray                                                      |
| JsonEncode                  | write json to IOWriter                                                              |
| JsonMarshal                 | generate json bytes and return                                                      |

# Recommend
1. Create instance of MagicArray
```go
arr1 := ValueofStruct(UserInfo{UserId:123,UserName:"name",})
arr2 := ValueofStructs([]UserInfo{
   {UserId:11,UserName:"name1",},
   {UserId:22,UserName:"name2",}
})
arr3:= ValueOfSlice([]string{"123","456,"789""})

// Recommend use the above methods to make instance, that will be lesser to transform calculates.
// at last ,you can use the common method:
arr4 := Valuoef(map[string]any{
"column1":123,
"column2":true,
})

```
2. Iterator to loop the array, for example:
```go
    arr := array.ValueofStruct(IteratorInfo{
        Field1: "field1",
        Field2: 2,
        Field3: true,
    })
    iter := arr.Iter()
    for k, v := iter.FirstKV(); k != nil; k, v = iter.NextKV() {
        fmt.Printf("%s=%s\n", k.String(), v.String())
    }
```
output:
```
Field1=field1
Field2=2
Field3=true
```
3. JsonMarshal function to generate json, omitEmpty and convert to hump naming rule easily.
```go
users, _ := array.Valueof([]map[string]any{
		{
			"id":        1,
			"user_name": "bobby",
			"Age":       nil,
		}
	})
	bs, _ := array.JsonMarshal(
		users,
		array.JosnOptOmitEmpty(true),
		array.JsonOptDefaultNamingUnderscoreToHump())

	fmt.Println(string(bs))
```
output:
````
[{"id":1,"userName":"bobby"}]

````

# Hello world

```go

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

func main() {

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

	mix, _ := array.Merge(array.ValueofStruct(user), score, area)
	mix = array.Pick(mix, "Id", "City", "Score")
	if bs, err := json.Marshal(mix); err == nil {
		fmt.Println(string(bs))
	} else {
		panic(err)
	}

}

```
output:
```json
{"userid":1,"City":"beij","Score":66}
```

# examples
More of examples, visit [_examples/](https://github.com/lingdor/magicarray/tree/main/_examples) in my repository. 

thanks!






