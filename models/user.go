package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"fullname" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255);unique;not_null"`
	Role      string    `json:"role" gorm:"type: varchar(255)"`
	Phone     int       `json:"phone" gorm:"type:varchar(255)"`
	Address   string    `json:"address" gorm:"type:varchar(255)"`
	Gender    string    `json:"gender" gorm:"type:varchar(255)"`
	Image     string    `json:"image" gorm:"type:varchar(255)"`
	Password  string    `json:"password" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
