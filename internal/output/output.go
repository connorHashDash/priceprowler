package output

import (
	"fmt"
	"math"
	"priceprowler/internal/hmlandreg"
	"sort"

	"github.com/fatih/color"
)

func TrendByPropertyType() error {
	TrendData, err := hmlandreg.GetPriceChange_AllTypes()
	if err != nil {
		return err
	}

	var PropertyType = []byte{
		'D',
		'S',
		'T',
		'F',
		'O',
	}

	var months = make(map[string]map[byte]int)

	for _, r := range TrendData {
		if months[r.Month] == nil {
			months[r.Month] = make(map[byte]int)
		}
		months[r.Month][r.PropertyType[0]] = r.AvgPrice
	}
	sortedMonths := make([]string, 0, len(TrendData))

	for month := range months {
		sortedMonths = append(sortedMonths, month)
	}

	sort.Strings(sortedMonths)

	fmt.Printf("Month\tDetac\tSemi\tTerrac\tFlat\tOther")
	prevMonth := ""
	for _, month := range sortedMonths {
		fmt.Printf("\n%v\t", month)
		for _, p := range PropertyType {
			var arrow rune
			var colourFunc func(format string, a ...any)
			price := months[month][p]
			fmt.Printf("%d", price)
			if prevMonth != "" {
				var diff float64 = calculatePercentDiff(float64(months[prevMonth][p]), float64(months[month][p]))
				colourFunc, arrow = colourOutput(diff)
				colourFunc("%v\t", string(arrow))
			}
		}
		prevMonth = month
	}

	return nil
}

func WholePostCodeTrend() error {
	data, err := hmlandreg.GetPriceChange_WholePostcode()
	if err != nil {
		return err
	}

	fmt.Printf("\n\n")
	prevVal := 0
	for k, v := range data {
		if k == 0 {
			fmt.Println("---Whole Area---")
		}
		percentage := calculatePercentDiff(float64(prevVal), float64(v.AvgPrice))
		colourFunc, arrow := colourOutput(percentage)

		fmt.Printf("%v\t", v.Month)
		if math.IsInf(percentage, 1) {
			colourFunc("%v %v\n", v.AvgPrice, string(arrow))
		} else {
			colourFunc("%v %6.2f%% %v\n", v.AvgPrice, percentage, string(arrow))
		}
		prevVal = v.AvgPrice

		// Final Difference Calculation
		if k == len(data)-1 {
			var firstValue int = data[0].AvgPrice
			var lastValue int = data[len(data)-1].AvgPrice

			var totalDiff float64 = calculatePercentDiff(float64(firstValue), float64(lastValue))
			colourFunc, arrow = colourOutput(totalDiff)
			fmt.Printf("Total change in postcode: ")
			colourFunc("%6.2f%% %v\n", totalDiff, string(arrow))
		}
	}

	return nil
}

func colourOutput(percentage float64) (func(format string, a ...any), rune) {
	var colourFunc func(format string, a ...any)
	var arrow rune
	if math.IsInf(percentage, 1) {
		colourFunc = color.New(color.FgWhite).PrintfFunc()
		arrow = ' '
	} else if percentage < 0 {
		colourFunc = color.New(color.FgRed).PrintfFunc()
		arrow = '↓'
	} else if percentage > 0 {
		colourFunc = color.New(color.FgGreen).PrintfFunc()
		arrow = '↑'
	} else {
		colourFunc = color.New(color.FgWhite).PrintfFunc()
		arrow = ' '
	}

	return colourFunc, arrow
}

func calculatePercentDiff(PriceA, PriceB float64) float64 {

	return (100 - ((PriceB / PriceA) * 100)) * -1
}
