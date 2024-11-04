package models

import "time"

type ExchangeRate struct {
	ID           uint      `gorm:"primarykey" json:"_id"`
	FromCurrency string    `json:"fromCurrency"`
	ToCurrency   string    `json:"toCurrency"`
	Rate         float64   `json:"rate" binding:"required"`
	Date         time.Time `json:"date"`
}
