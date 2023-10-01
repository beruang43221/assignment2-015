package models

import "time"

type Item struct {
	Item_id uint `gorm:"primaryKey" json:"item_id"`
	Item_code string `json:"item_code"`
	Quantity uint `json:"quantity"`
	Description string `json:"description"`
	Order_id uint `json:"order_id"`
	Created_at time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}