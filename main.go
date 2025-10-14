package main

import (
	"log"
	"priceprowler/internal/hmlandreg"
	"priceprowler/internal/output"
)

func main() {
	hmlandreg.Init()
	err := output.TrendByPropertyType()
	if err != nil {
		log.Fatal(err)
	}
	output.WholePostCodeTrend()
}
