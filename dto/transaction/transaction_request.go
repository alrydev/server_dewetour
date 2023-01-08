package transactionsdto

type CreateTransactionRequest struct {
	CounterQty int    `json:"counter_qty" form:"counter_qty" gorm:"type: varchar(255)"`
	UserId     int    `json:"user_id" form:"userId" gorm:"type: varchar(255)"`
	Total      int    `json:"total" form:"total" gorm:"type: varchar(255)"`
	Status     string `json:"status" form:"status" gorm:"type: varchar(255)"`
	// Attachment string `json:"-" form:"image" gorm:"type: varchar(255)"`
	TripId int `json:"trip_id" form:"tripId" gorm:"type: varchar(255)"`
}

type UpdateTransactionRequest struct {
	CounterQty int    `json:"counter_qty" form:"counter_qty" gorm:"type: varchar(255)"`
	UserId     int    `json:"userId" for:"userId" gorm:"type: varchar(255)"`
	Total      int    `json:"total" form:"total" gorm:"type: varchar(255)"`
	Status     string `json:"status" form:"status" gorm:"type: varchar(255)"`
	// Attachment string `json:"-" form:"image" gorm:"type: varchar(255)"`
	TripId int `json:"tripId" form:"tripId" gorm:"type: varchar(255)"`
}
