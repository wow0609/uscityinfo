package main

import (
	"log"

	"github.com/wow0609/uscityinfo/api"
	"github.com/wow0609/uscityinfo/dataAccess"
)

func main() {
	address := "localhost:3000"
	repo := dataAccess.NewDummyCityData()
	api := api.CreateApiServer(address, repo)
	log.Fatal(api.StartApiServer())
}
