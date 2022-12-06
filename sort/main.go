package main

import (
	"fmt"
	"sort"
	"time"
)

type Ell struct {
	Date time.Time
}

func main() {
	arr := []*Ell{
		{
			Date: time.Date(2022, time.April, 23, 0, 0, 0, 0, time.UTC),
		},
		{
			Date: time.Date(2022, time.April, 21, 0, 0, 0, 0, time.UTC),
		},
		{
			Date: time.Date(2022, time.April, 20, 0, 0, 0, 0, time.UTC),
		},
		{
			Date: time.Date(2022, time.April, 22, 0, 0, 0, 0, time.UTC),
		},
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Date.After(arr[j].Date)
	})
	for _, v := range arr {
		fmt.Println(v.Date)
	}
}
