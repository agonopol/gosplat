package main

import (
	"gosplat"
)

func Timeseries() {
	f := gosplat.NewFrame("Timeseires Stick Example Frame")
	row := gosplat.NewRow("Timeserieseses")
	v1 := gosplat.NewChart()
	v1.Append(map[string]interface{}{"x": 100, "thing": 20, "thong": 11})
	v1.Append(map[string]interface{}{"x": 200, "thing": 24, "thong": 19})
	v1.Append(map[string]interface{}{"x": 300, "thing": 24, "thong": 19})
	v1.Append(map[string]interface{}{"x": 400, "thing": 14})
	v1.Append(map[string]interface{}{"x": 500, "thing": 17})
	v1.Append(map[string]interface{}{"x": 600, "thong": -3})
	row.Append("Bucketed time series", v1.Timeseries())

	v2 := gosplat.NewChart()
	v2.Append(map[string]interface{}{"x": "2011/07/24", "thing": 20, "thong": 11})
	v2.Append(map[string]interface{}{"x": "2011/07/25", "thing": 24, "thong": 19})
	v2.Append(map[string]interface{}{"x": "2011/07/26", "thing": 24, "thong": 19})
	row.Append("Time series", v2.Timeseries())

	f.AppendRow(row)

	err := f.Preview()

	if err != nil {
		panic(err)
	}
}
