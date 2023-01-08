package tripsdto

type TripResponse struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	CountryID      int    `json:"country_id"`
	Accomodation   string `json:"accomodation"`
	Transportation string `json:"transportation"`
	Meal           string `json:"meal"`
	Day            int    `json:"day"`
	Night          int    `json:"night"`
	DateTrip       string `json:"date"`
	Price          int    `json:"price"`
	Quota          int    `json:"quota"`
	Desc           string `json:"desc"`
	Image          string `json:"image"`
}
