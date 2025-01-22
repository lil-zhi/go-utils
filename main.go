package main

import (
	"fmt"
	calcutil "github.com/lil-zhi/go-utils/calc-util"
)

type Part struct {
	PartID   int64
	PartName string
}

type B struct {
	UID  int64
	Part Part
}

type A struct {
	ID   int64
	Age  int64
	Part Part
}

func main() {
	var a = []A{
		{ID: 1, Age: 1},
		{ID: 2, Age: 2},
		{ID: 3, Age: 3},
		{ID: 4, Age: 4},
	}
	var a1 = []A{
		{ID: 1, Age: 1},
		{ID: 2, Age: 2},
		{ID: 3, Age: 3},
		{ID: 5, Age: 5},
		{ID: 6, Age: 6},
	}
	fmt.Println(calcutil.Com(a, a1))
	fmt.Println(calcutil.Sub(a, a1))
	var b = []B{
		{UID: 1, Part: Part{
			PartID:   1,
			PartName: "1",
		}}, {UID: 2, Part: Part{
			PartID:   2,
			PartName: "2",
		}}, {UID: 3, Part: Part{
			PartID:   3,
			PartName: "3",
		}}, {UID: 4, Part: Part{
			PartID:   4,
			PartName: "4",
		}},
	}
	res := calcutil.UpdateListWithList(a, b, func(a A) int64 {
		return a.ID
	}, func(b B) int64 {
		return b.UID
	}, func(a A, b B) A {
		a.Part = b.Part
		return a
	})
	fmt.Println(res)
	var c = []B{
		{UID: 1, Part: Part{
			PartID:   1,
			PartName: "1",
		}}, {UID: 3, Part: Part{
			PartID:   3,
			PartName: "3",
		}},
	}
	res = calcutil.FilterListByList(a, c, func(a A) int64 {
		return a.ID
	}, func(b B) int64 {
		return b.UID
	})
	fmt.Println(res)
}
