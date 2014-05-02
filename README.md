gosplat
=======

Easily generate html/js graphs in go with dygraphs/gocharts etc...

## Installation:

go get github.com/agonopol/gosplat

## Import:

import "github.com/agonopol/gosplat"

## Example:
```golang
func Linechart() {
	f := gosplat.NewFrame("Linechart Example Frame")
	v := gosplat.NewChart()
	v.Append(map[string]interface{}{"date": "2011/07/23", "thing": 10, "thong": 20, "whatevs": 14})
	v.Append(map[string]interface{}{"date": "2011/07/24", "thing": 12, "thong": 24, "whatevs": 24})
	v.Append(map[string]interface{}{"date": "2011/07/24", "thing": 12, "thong": 7, "whatevs": 11})
	f.Append("Linechart Example", v.Linechart())
	err := f.Preview()
	if err != nil {
		panic(err)
	}
}
```
## Supported Charts:

-barchart
-candlestick
-columnchart
-combochart
-linechart
-piechart
-scatter
-table
-timeseries



