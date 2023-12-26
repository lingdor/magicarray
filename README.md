# magicarray
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

# Hello world

```go

package main

import (
	"encoding/json"
	"fmt"
	arr "github.com/lingdor/magicarray"
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

	mix, _ := arr.Merge(arr.ValueofStruct(user), score, area)
	mix = arr.Pick(mix, "Id", "City", "Score")
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






