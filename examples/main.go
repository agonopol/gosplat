package main

import (
	"flag"
	"os"
)

var plot *string = flag.String("plot", "", `Plot to Preview.  Available:
	barchart
	candlestick
	columnchart
	combochart
	linechart
	piechart
	scatter
	table
	timeseries
	`)
var help *bool = flag.Bool("help", false, "Show help")

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(1)
	}

	switch *plot {
	case "barchart":
		Barchart()
		break
	case "candlestick":
		Candlestick()
		break
	case "columnchart":
		Columnchart()
		break
	case "combochart":
		Combochart()
		break
	case "linechart":
		Linechart()
		break
	case "piechart":
		Piechart()
		break
	case "scatter":
		Scatter()
		break
	case "table":
		Table()
		break
	case "timeseries":
		Timeseries()
		break
	default:
		flag.Usage()
		os.Exit(1)
		break
	}
}
