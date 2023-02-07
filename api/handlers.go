package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (a *ApiServer) registerHandlers() {
	log.Printf("==> Registering handlers for Api Server ... ")

	//http.HandleFunc("/api/{cityName}", a.cityHandler)
	http.HandleFunc("/api/", a.rootApiHandler)

	log.Printf("==> Registering handlers [DONE] ")
}

func (a *ApiServer) rootApiHandler(w http.ResponseWriter, r *http.Request) {
	cityName := r.URL.Path[5:]
	log.Printf("Parameter for cityName: %s", cityName)

	cr := a.repo.CityByName(cityName)

	json.NewEncoder(w).Encode(cr)
}

//func (a *ApiServer) cityHandler(w http.ResponseWriter, r *http.Request) {
//	//w.Write()"ApiServer Root Handler")
//}
