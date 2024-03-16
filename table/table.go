package table

import (
	"juliano.com/passage-order/mapp"
	"juliano.com/passage-order/theme"
)

var CellStyle = map[string]string{
	"padding":          "17px 42px 17px 42px",
	"text-align":       "left",
	"color":            "#303030",
	"background-color": "white",
	"font-size":        "0.92rem",
}

var ThStyle = mapp.Merge(CellStyle, map[string]string{
	// "background-color": "#f2f2f2",
	// "background-color": "#e0f2fe", // Previous value
	// "background-color": "#bae6fd",
	"background-color": "rgb(112, 123, 211)",
	"color":            "white",
	"font-size":        "1rem",
})

var OddTdStyle = mapp.Merge(CellStyle, map[string]string{
	"background-color": "rgb(248, 247, 254)",
})

var TableStyle = map[string]string{
	"border-collapse": "collapse",
	// "box-shadow":      "rgba(100, 100, 111, 0.2) 0px 7px 29px 0px",
	"box-shadow":    "rgba(99, 99, 99, 0.2) 0px 2px 8px 0px",
	"border-radius": theme.SmallRadius.String(),
	"overflow":      "hidden",
}

func TdStyle(index int) map[string]string {
	if index%2 == 1 {
		return OddTdStyle
	}

	return CellStyle
}

func IconStyle(index int) map[string]string {
	style := TdStyle(index)

	return mapp.Merge(style, map[string]string{
		"cursor": "pointer",
	})
}

func IconStyleFromStyle(style map[string]string) map[string]string {
	return mapp.Merge(style, map[string]string{
		"cursor": "pointer",
	})
}
