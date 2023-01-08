package handlers

import (
	"context"
	dto "dewe/dto/result"
	tripsdto "dewe/dto/trips"
	"dewe/models"
	"dewe/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerTrips struct {
	TripRepository repositories.TripsRepository
}

func HandlerTrips(TripRepository repositories.TripsRepository) *handlerTrips {
	return &handlerTrips{TripRepository}
}

func (h *handlerTrips) FindTrips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	trips, err := h.TripRepository.FindTrips()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, p := range trips {
		trips[i].Image = p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trips}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrips) GetTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trip}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerTrips) CreateTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	country_id, _ := strconv.Atoi(r.FormValue("country_id"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	quota, _ := strconv.Atoi(r.FormValue("quota"))
	day, _ := strconv.Atoi(r.FormValue("day"))
	night, _ := strconv.Atoi(r.FormValue("night"))
	request := tripsdto.CreateTripRequest{
		Title:          r.FormValue("title"),
		CountryID:      country_id,
		Accomodation:   r.FormValue("accomodation"),
		Transportation: r.FormValue("transportation"),
		Meal:           r.FormValue("meal"),
		Day:            day,
		Night:          night,
		DateTrip:       r.FormValue("date"),
		Price:          price,
		Quota:          quota,
		Desc:           r.FormValue("desc"),
		Image:          r.FormValue("image"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "dewetrips"})
	if err != nil {
		fmt.Println(err.Error())
	}

	trip := models.Trip{
		Title:          request.Title,
		CountryID:      request.CountryID,
		Accomodation:   request.Accomodation,
		Transportation: request.Transportation,
		Meal:           request.Meal,
		Day:            request.Day,
		Night:          request.Night,
		Date:           request.DateTrip,
		Price:          request.Price,
		Quota:          request.Quota,
		Desc:           request.Desc,
		Image:          resp.SecureURL,
	}

	data, err := h.TripRepository.CreateTrip(trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrips) UpdateTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// request := new(tripsdto.UpdateTripRequest)
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "dewetrips"})
	if err != nil {
		fmt.Println(err.Error())
	}

	country_id, _ := strconv.Atoi(r.FormValue("country_id"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	quota, _ := strconv.Atoi(r.FormValue("quota"))
	day, _ := strconv.Atoi(r.FormValue("day"))
	night, _ := strconv.Atoi(r.FormValue("night"))
	request := tripsdto.UpdateTripRequest{
		Title:          r.FormValue("title"),
		CountryID:      country_id,
		Accomodation:   r.FormValue("accomodation"),
		Transportation: r.FormValue("transportation"),
		Meal:           r.FormValue("meal"),
		Day:            day,
		Night:          night,
		DateTrip:       r.FormValue("date"),
		Price:          price,
		Quota:          quota,
		Desc:           r.FormValue("desc"),
		Image:          resp.SecureURL,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	trip, err := h.TripRepository.GetTrip(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		trip.Title = request.Title
	}
	if request.CountryID != 0 {
		trip.CountryID = request.CountryID
	}
	if request.Accomodation != "" {
		trip.Accomodation = request.Accomodation
	}
	if request.Transportation != "" {
		trip.Transportation = request.Transportation
	}
	if request.Meal != "" {
		trip.Meal = request.Meal
	}
	if request.Day != 0 {
		trip.Day = request.Day
	}
	if request.Night != 0 {
		trip.Night = request.Night
	}
	if request.DateTrip != "" {
		trip.Date = request.DateTrip
	}
	if request.Price != 0 {
		trip.Price = request.Price
	}
	if request.Quota != 0 {
		trip.Quota = request.Quota
	}
	if request.Desc != "" {
		trip.Desc = request.Desc
	}
	if request.Image != "" {
		trip.Image = filepath
	}

	data, err := h.TripRepository.UpdateTrip(trip, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerTrips) DeleteTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatio/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TripRepository.DeleteTrip(trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	responses := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(responses)
}

func convertResponseTrip(u models.Trip) tripsdto.TripResponse {
	return tripsdto.TripResponse{
		ID:             u.ID,
		Title:          u.Title,
		CountryID:      u.CountryID,
		Accomodation:   u.Accomodation,
		Transportation: u.Transportation,
		Meal:           u.Meal,
		Day:            u.Day,
		Night:          u.Night,
		DateTrip:       u.Date,
		Price:          u.Price,
		Quota:          u.Quota,
		Desc:           u.Desc,
		Image:          u.Image,
	}

}
