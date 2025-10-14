package hmlandreg

import "time"

type houseSales struct {
	TransactionId string    `json:"transaction_id"`
	Price         int       `json:"price"`
	TransferDate  time.Time `json:"transfer_date"`
	PostCode      string    `json:"postcode"`
	PropertyType  string    `json:"property_type"`
	NewBuild      rune      `json:"new_build"`
	Tenure        rune      `json:"tenure"`
	Paon          string    `json:"paon"`
	Sreet         string    `json:"street"`
	RecordStatus  rune      `json:"record_status"`
}

type PriceTrendData struct {
	Month        string
	PropertyType []byte
	AvgPrice     int
}

type WholePostCodeTrend struct {
	Month    string
	AvgPrice int
}
