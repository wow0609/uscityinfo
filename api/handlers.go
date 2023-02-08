package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *ApiServer) registerHandlers() {
	log.Printf("==> Registering handlers for Api Server ... ")

	a.router.HandleFunc("/api/city/{cityName}/{year}", a.cityYearHandler)
	a.router.HandleFunc("/api/city/{cityName}", a.cityHandler)
	a.router.HandleFunc("/api/city", a.allCitiesHandler).Methods("GET")

	http.Handle("/api/", a.router)

	log.Printf("==> Registering handlers [DONE] ")
}

func (a *ApiServer) allCitiesHandler(w http.ResponseWriter, r *http.Request) {

	cr := a.cityRepo.AllCities()

	json.NewEncoder(w).Encode(cr)
}

func (a *ApiServer) cityHandler(w http.ResponseWriter, r *http.Request) {
	cityName := mux.Vars(r)["cityName"]
	if len(cityName) > 0 {
		result := a.cityRepo.CityByName(cityName)
		if result != nil {
			json.NewEncoder(w).Encode(result)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func (a *ApiServer) cityYearHandler(w http.ResponseWriter, r *http.Request) {
	//log.Print("In the cityYear Handler...")
	cityName := mux.Vars(r)["cityName"]
	sYear := mux.Vars(r)["year"]
	if len(cityName) > 0 {
		result := a.cityRepo.CityByName(cityName)
		if result != nil {
			iYear, err := strconv.Atoi(sYear)
			if err != nil {
				http.Error(w, "Invalid Request:  Cannot parse Year.", http.StatusBadRequest)
				return
			}

			demoData := result.FindDemographicsByYear(iYear)
			if demoData == nil {
				http.Error(w, "Year not found for City.", http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(demoData)
		} else {
			http.Error(w, "City not found.", http.StatusNotFound)
			//w.WriteHeader(http.StatusNotFound)
		}
	}
}
