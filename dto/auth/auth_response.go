package authdto

type LoginResponse struct {
	Name  string `gorm:"type: varchar(255)" json:"Name"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	// Role     string `gorm:"type: varchar(50)"  json:"role"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}

type RegisterResponse struct {
	Name     string `gorm:"type: varchar(255)" json:"Name"`
	Password string `gorm:"type: varchar(255)" json:"password"`
}

type CheckAuthResponse struct {
	Id      int    `gorm:"type: int" json:"id"`
	Name    string `gorm:"type: varchar(255)" json:"Name"`
	Email   string `gorm:"type: varchar(255)" json:"email"`
	Image   string `gorm:"type: varchar (255)" json:"image"`
	Phone   int    `gorm:"type: varchar(255)" json:"phone"`
	Gender  string `gorm:"type: varchar(255)" json:"gender"`
	Address string `gorom:"type: varchar(255)" json:"address"`
	Role    string `gorm:"type: varchar(50)"  json:"role"`
}
