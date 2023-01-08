package models

import "time"

type Transaction struct {
	ID         int       `json:"id"`
	UserID     int       `json:"userId"`
	User       User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CounterQty int       `json:"counter_qty" gorm:"type:varchar(255)"`
	Total      int       `json:"total" gorm:"type:varchar(255)"`
	Status     string    `json:"status" gorm:"type:varchar(255)"`
	TripId     int       `json:"trip_id"`
	Trip       Trip      `json:"trip" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt  time.Time `json:"created_at"`
}
