package main

import (
	"fmt"
	arr "github.com/lingdor/magicarray"
)

type IteratorInfo struct {
	Field1 string
	Field2 int
	Field3 bool
}

func iteratorCommand() {
	ma := arr.ValueofStruct(IteratorInfo{
		Field1: "field1",
		Field2: 2,
		Field3: true,
	})
	iter := ma.Iter()
	for k, v := iter.FirstKV(); k != nil; k, v = iter.NextKV() {
		fmt.Printf("%s=%s\n", k.String(), v.String())
	}

}
