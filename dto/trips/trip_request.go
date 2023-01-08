package tripsdto

type CreateTripRequest struct {
	Title          string `json:"title" form:"title"`
	CountryID      int    `json:"country_id" form:"country_id"`
	Accomodation   string `json:"accomodation" form:"accomodation"`
	Transportation string `json:"transportation" form:"transportation"`
	Meal           string `json:"meal" form:"meal"`
	Day            int    `json:"day" form:"day"`
	Night          int    `json:"night" form:"night"`
	DateTrip       string `json:"date" form:"date"`
	Price          int    `json:"price" form:"price"`
	Quota          int    `json:"quota" form:"quota"`
	Desc           string `json:"desc" form:"desc"`
	Image          string `json:"image" form:"image"`
}

type UpdateTripRequest struct {
	Title          string `json:"title" form:"title"`
	CountryID      int    `json:"country_id" form:"country_id"`
	Accomodation   string `json:"accomodation" form:"accomodation"`
	Transportation string `json:"transportation" form:"transportation"`
	Meal           string `json:"meal" form:"meal"`
	Day            int    `json:"day" form:"day"`
	Night          int    `json:"night" form:"night"`
	DateTrip       string `json:"date" form:"date"`
	Price          int    `json:"price" form:"price"`
	Quota          int    `json:"quota" form:"quota"`
	Desc           string `json:"desc" form:"desc"`
	Image          string `json:"image" form:"image"`
}
