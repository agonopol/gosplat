package main

import (
	"github.com/agonopol/gosplat"
)

func Piechart() {
	f := gosplat.NewFrame("Piechart Example Frame")
	v := gosplat.NewChart()
	v.Append(map[string]interface{}{"task": "sleep", "hours": 10})
	v.Append(map[string]interface{}{"task": "rock", "hours": 12})
	v.Append(map[string]interface{}{"task": "tv", "hours": 2})
	f.Append("Pie chart", v.Piechart())

	err := f.Preview()
	if err != nil {
		panic(err)
	}
}
