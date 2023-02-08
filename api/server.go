package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wow0609/uscityinfo/dataAccess"
)

type ApiServer struct {
	address  string
	router   *mux.Router
	cityRepo dataAccess.CityRepository
}

func CreateApiServer(serverAddress string, dataAccess dataAccess.CityRepository) *ApiServer {
	return &ApiServer{
		address:  serverAddress,
		cityRepo: dataAccess,
		router:   mux.NewRouter(),
	}
}

func (a *ApiServer) StartApiServer() error {
	log.Print("Starting Api Server")
	a.registerHandlers()
	log.Printf("Api Server listening on [%s]... ", a.address)
	return http.ListenAndServe(a.address, nil)
}
