package countrydto

type CreateCountryRequest struct {
	Name string `json:"name" form:"name" gorm:"type: varchar(255)"`
}

type UpdateCountryRequest struct {
	Name string `json:"name" form:"name" gorm:"type: varchar(255)"`
}
