gosplat
=======

Easily generate html/js graphs in go with dygraphs/gocharts etc...

gosplat generates an html file with your and js embded in the page to plot the data.
## Installation:

`go get github.com/agonopol/gosplat`

## Import:

`import "github.com/agonopol/gosplat"`

## Example:

<img src="https://raw.githubusercontent.com/agonopol/gosplat/master/examples/example.png"  width="673" height="287"/>

```go
func Linechart() {
	//Create a frame to put charts in
	f := gosplat.NewFrame("Linechart Example Frame")
	
	//Create a chart
	v := gosplat.NewChart()
	
	//Add some random data
	v.Append(map[string]interface{}{"date": "2011/07/23", "thing": 10, "thong": 20, "whatevs": 14})
	v.Append(map[string]interface{}{"date": "2011/07/24", "thing": 12, "thong": 24, "whatevs": 24})
	v.Append(map[string]interface{}{"date": "2011/07/24", "thing": 12, "thong": 7, "whatevs": 11})
	
	//Add the chart to the Frame
	f.Append("Linechart Example", v.Linechart())
	
	//Preview generates a tmp html file and opens it with the default browser
	err := f.Preview()
	if err != nil {
		panic(err)
	}
}
```


## Supported Charts:

* barchart
* candlestick
* columnchart
* combochart
* linechart
* piechart
* scatter
* table
* timeseries



