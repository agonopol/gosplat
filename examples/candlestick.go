package main

import (
	"github.com/agonopol/gosplat"
)

func Candlestick() {
	f := gosplat.NewFrame("Candlestick Example Frame")
	v := gosplat.NewChart()
	v.Append([]interface{}{"rock", 20, 28, 38, 45})
	v.Append([]interface{}{"metal", 31, 38, 55, 66})
	v.Append([]interface{}{"blues", 50, 55, 77, 80})
	f.Append("Candlestick Example", v.Candlestick())
	err := f.Preview()
	if err != nil {
		panic(err)
	}
}
