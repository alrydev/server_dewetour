package models

import "time"

type Trip struct {
	ID             int       `json:"id" gorm:"primary_key:auto_increment"`
	Title          string    `json:"title" gorm:"type:varchar(255)"`
	CountryID      int       `json:"country_id"`
	Country        Country   `json:"country" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Accomodation   string    `json:"accomodation" gorm:"type:varchar(255)"`
	Transportation string    `json:"transportation" gorm:"type:varchar(255)"`
	Meal           string    `json:"meal" gorm:"type:varchar(255)"`
	Day            int       `json:"day" gorm:"type:varchar(255)"`
	Night          int       `json:"night" gorm:"type:varchar(255)"`
	Date           string    `json:"dateTrip" gorm:"type:varchar(255)"`
	Price          int       `json:"priceTrip" gorm:"type:varchar(255)"`
	Quota          int       `json:"quota" gorm:"type:varchar(255)"`
	Desc           string    `json:"desc" gorm:"type:varchar(255)"`
	Image          string    `json:"image" gorm:"type:varchar(255)"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
