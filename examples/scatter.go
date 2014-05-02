package main

import (
	"gosplat"
)

func Scatter() {
	f := gosplat.NewFrame("Scatter Stick Example Frame")
	v := gosplat.NewChart()
	v.Append(map[string]interface{}{"x": 100, "thing": 20, "thong": 11})
	v.Append(map[string]interface{}{"x": 200, "thing": 24, "thong": 19})
	v.Append(map[string]interface{}{"x": 300, "thing": 24, "thong": 19})
	v.Append(map[string]interface{}{"x": 100, "thang": 14})
	v.Append(map[string]interface{}{"x": 100, "thung": 17})
	v.Append(map[string]interface{}{"x": 200, "thung": -3})
	f.Append("Scatter Example", v.Scatter())
	err := f.Preview()
	if err != nil {
		panic(err)
	}
}
