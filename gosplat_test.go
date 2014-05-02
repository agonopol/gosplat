package gosplat

import (
	"testing"
)


func TestPrintHtml(t * testing.T) {
	f := NewFrame("Test Frame")
	v := NewVisualize()
	v.Append(map[string]interface{}{"x":100, "y":200})
	v.Append(map[string]interface{}{"x":100, "y":300})
	v.Append(map[string]interface{}{"x":150, "y":200})
	f.Append("Scatter Example", v.Scatter())
	err := f.Preview()
	if (err != nil) {
		panic(err)
	}
}