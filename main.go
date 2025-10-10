package main

import (
	"fmt"
	"priceprowler/internal/hmlandreg"
)

func main() {
	hmlandreg.Init()
	TrendData, err := hmlandreg.GetPriceChange_AllTypes()
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range TrendData {
		fmt.Println(v.Month, string(v.PropertyType), v.AvgPrice)
	}

}
