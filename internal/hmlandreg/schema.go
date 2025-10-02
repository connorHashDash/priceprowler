package hmlandreg

import "time"

type houseSales struct {
	TransactionId string    `json:"transaction_id" gorm:"column:transaction_id"`
	Price         int       `json:"price" gorm:"column:price"`
	TransferDate  time.Time `json:"transfer_date" gorm:"column:transfer_date"`
	PostCode      string    `json:"postcode" gorm:"column:postcode"`
	PropertyType  string    `json:"property_type" gorm:"column:property_type"`
	NewBuild      rune      `json:"new_build" gorm:"column:new_build"`
	Tenure        rune      `json:"tenure" gorm:"column:tenure"`
	Paon          string    `json:"paon" gorm:"column:paon"`
	Sreet         string    `json:"street" gorm:"column:street"`
	RecordStatus  rune      `json:"record_status" gorm:"column:record_status"`
}
