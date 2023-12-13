# magicarray
NO care type, no care struct, easy to sort and aggregation.similar with php array structure,easy to deal for data process. Coding shortly,process automatically.
in magicarray, no nil forever.
```shell
go get github.com/lingdor/magicarray
```
```go
package main

import (
	"fmt"
	arr "github.com/lingdor/magicarray"
)

type struct1 struct {
	Field1 string
	Field2 int
	Field3 bool
}

func main() {

	rows, _ := arr.Valueof([]struct1{
		{Field1: "row1", Field2: 1, Field3: true},
		{Field1: "row2", Field2: 2, Field3: true},
		{Field1: "row3", Field2: 3, Field3: true},
	})

	fmt.Println("column:")
	iter := arr.Column(rows, "Field2").Iter()
	for val := iter.FirstVal(); val != nil; val = iter.NextVal() {
		fmt.Println(val.MustInt())
	}
	fmt.Println("")

	fmt.Println("dto translate:")

	entity, _ := arr.Valueof(struct1{Field1: "row1", Field2: 1, Field3: true})

	userInfo := map[string]any{
		"userid":   1,
		"username": "tom",
		"age":      88,
	}

	mixDto, _ := arr.Merge(entity, userInfo)

	if bs, err := arr.ToJson(mixDto); err == nil {
		fmt.Println(string(bs))
	} else {
		panic(err)
	}
}

```

output:
column:
1
2
3

dto translate:
{"Field1":"row1","Field2":1,"Field3":true,"age":88,"userid":1,"username":"tom"}





