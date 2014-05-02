package main

import (
	"gosplat"
)

func Columnchart() {
	f := gosplat.NewFrame("Columnchart Example Frame")
	v := gosplat.NewVisualize()
	v.Append(map[string]interface{}{"date": "2011/07/23", "thing": 10, "thong": 20, "whatevs": 14})
	v.Append(map[string]interface{}{"date": "2011/07/24", "thing": 12, "thong": 24, "whatevs": 24})
	v.Append(map[string]interface{}{"date": "2011/07/24", "thing": 12, "thong": 7, "whatevs": 11})
	f.Append("Columnchart Example", v.Columnchart())
	err := f.Preview()
	if err != nil {
		panic(err)
	}
}
