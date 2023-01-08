package authdto

type RegisterRequest struct {
	Name     string `gorm:"type: varchar(255)" json:"name" form:"name" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"email" form:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" form:"password" validate:"required"`
	Phone    int    `json:"phone" gorm:"type: varchar(255)" form:"phone" validate:"required"`
	Address  string `json:"address" gorm:"type: varchar(255)" form:"address" validate:"required"`
	Gender   string `json:"gender" gorm:"type: varchar(255)" form:"gender" validate:"required"`
	Image    string `json:"image" gorm:"type: varchar(255)" form:"image"`
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}
