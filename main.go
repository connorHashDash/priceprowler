package main

import (
	"fmt"
	"priceprowler/internal/hmlandreg"

	"github.com/fatih/color"
)

func main() {
	hmlandreg.Init()
	TrendData, err := hmlandreg.GetPriceChange_AllTypes()
	if err != nil {
		fmt.Println(err)
	}

	var wholeTypes = map[byte]string{
		'D': "Detached",
		'S': "Semi-Detached",
		'T': "Terraced",
		'F': "Flat",
		'O': "Other",
	}

	var prevType string
	var prevPrice int
	for _, v := range TrendData {
		if prevType != string(v.PropertyType) {
			fmt.Printf("\n")
			color.Set(color.Bold)
			fmt.Println("---" + wholeTypes[v.PropertyType[0]] + "---")
			color.Unset()
			prevPrice = 0
		}
		var priceColor func(format string, a ...any)
		if prevPrice > v.AvgPrice && prevPrice > 0 {
			priceColor = color.New(color.FgRed).PrintfFunc()
		} else if prevPrice < v.AvgPrice && prevPrice > 0 {
			priceColor = color.New(color.FgGreen).PrintfFunc()
		} else if prevPrice == 0 {
			priceColor = color.New(color.FgWhite).PrintfFunc()
		}
		prevPrice = v.AvgPrice

		prevType = string(v.PropertyType)
		fmt.Printf("%v\t", v.Month)
		priceColor("%v\n", v.AvgPrice)
	}
}
