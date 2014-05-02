package main

import (
	"github.com/agonopol/gosplat"
)

func Combochart() {
	f := gosplat.NewFrame("Combochart Stick Example Frame")
	v := gosplat.NewChart()
	v.Append(map[string]interface{}{"Month": "2004/05", "Bolivia": 165, "Ecuador": 938, "Average": (165 + 938) / 2})
	v.Append(map[string]interface{}{"Month": "2004/06", "Bolivia": 135, "Ecuador": 1120, "Average": (135 + 1120) / 2})
	v.Append(map[string]interface{}{"Month": "2004/07", "Bolivia": 157, "Ecuador": 1167, "Average": (157 + 1167) / 2})
	v.Append(map[string]interface{}{"Month": "2004/08", "Bolivia": 139, "Ecuador": 1110, "Average": (139 + 1110) / 2})
	v.Append(map[string]interface{}{"Month": "2004/09", "Bolivia": 136, "Ecuador": 691, "Average": (136 + 691) / 2})

	f.Append("Combochart Example", v.Combochart(map[string]interface{}{
		"title":      "Monthly Coffee Production by Country",
		"vAxis":      map[string]interface{}{"title": "Cups"},
		"hAxis":      map[string]interface{}{"title": "Month"},
		"seriesType": "bars",
		"series":     map[string]interface{}{"2": map[string]interface{}{"type": "line"}}}))
	err := f.Preview()
	if err != nil {
		panic(err)
	}
}
