package array

import (
	"testing"
)

func TestColumn(t *testing.T) {

	type tt struct {
		Col1 string
		Col2 string
	}

	var var1 = []tt{
		{Col1: "good", Col2: "2222"},
		{Col1: "good_2", Col2: "2222_2"},
		{Col1: "good_3", Col2: "2222_3"},
	}
	expect := []string{
		"2222", "2222_2", "2222_3",
	}
	infos, err := Valueof(var1)
	if err != nil {
		t.Error(err)
	}

	column1 := Column(infos, "Col2")
	if err := Equals(column1, expect); err != nil {
		t.Error(err)
	}

}
