package main

import (
	"log"
	"os"

	"github.com/der-nackte-halloumi/api/datastore/postgres"
	"github.com/der-nackte-halloumi/api/interface/rest"
	"github.com/der-nackte-halloumi/api/usecase/search_shop"
)

func main() {

	datastore, err := postgres.NewDatastore(
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PASSWORD"),
	)

	if err != nil {
		log.Fatal(
			"could not open datastore connection: ",
			err,
		)
	}

	shopRepository := postgres.NewShopRepository(datastore)

	searchShopService := search_shop.NewService(shopRepository)

	apiPort := os.Getenv("API_PORT")
	restAPI, err := rest.NewRestAPI(searchShopService, apiPort)

	if err != nil {
		log.Fatal("server could not be started: ", err)
	}

	log.Printf("server started on port %s", apiPort)
	restAPI.Start()
}
