package model

import "time"

const (
	ORDER_STATUS_UNPAY    = 1
	ORDER_STATUS_PAYING   = 2
	ORDER_STATUS_PAID     = 3
	ORDER_STATUS_CANCELED = 4
)

type Order struct {
	Model
	UserId       uint       `json:"userId" gorm:"index"`
	Charge       int        `json:"charge"`
	Status       int        `json:"status"`
	IsRefund     int        `json:"isRefund"`
	LastRefundAt *time.Time `json:"lastRefundAt"`
}
