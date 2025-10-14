package output

import (
	"fmt"
	"github.com/fatih/color"
	"priceprowler/internal/hmlandreg"
)

func TrendByPropertyType() error {
	TrendData, err := hmlandreg.GetPriceChange_AllTypes()
	if err != nil {
		return err
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
	for k, v := range TrendData {
		if prevType != string(v.PropertyType) {
			if k != 0 {
				fmt.Printf("\n")
			}
			color.Set(color.Bold)
			fmt.Println("---" + wholeTypes[v.PropertyType[0]] + "---")
			color.Unset()
			prevPrice = 0
		}

		var priceColor func(format string, a ...any)
		var arrow rune
		if prevPrice > v.AvgPrice && prevPrice > 0 {
			priceColor = color.New(color.FgRed).PrintfFunc()
			arrow = '↓'
		} else if prevPrice < v.AvgPrice && prevPrice > 0 {
			priceColor = color.New(color.FgGreen).PrintfFunc()
			arrow = '↑'
		} else {
			priceColor = color.New(color.FgWhite).PrintfFunc()
			arrow = ' '
		}

		prevType = string(v.PropertyType)
		fmt.Printf("%v ", v.Month)
		priceColor("%v", v.AvgPrice)
		if prevPrice != 0 {
			priceColor("%6.2f%% ", calculatePercentDiff(float64(prevPrice), float64(v.AvgPrice)))
		}
		priceColor("%v", string(arrow))
		fmt.Printf("\n")
		prevPrice = v.AvgPrice

	}
	return nil
}

func WholePostCodeTrend() error {
	data, err := hmlandreg.GetPriceChange_WholePostcode()
	if err != nil {
		return err
	}

	for k, v := range data {
		fmt.Printf("%v,%v", k, v)
	}

	return nil
}

func calculatePercentDiff(PriceA, PriceB float64) float64 {

	return (100 - ((PriceB / PriceA) * 100)) * -1
}
