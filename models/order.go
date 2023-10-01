package models

import "time"

type Order struct {
	Order_id      uint `gorm:"primaryKey" json:"id"`
	Customer_name string `json:"customer_name"`
	Items []Item
	Ordered_at time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"ordered_at"`
	Created_at time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}