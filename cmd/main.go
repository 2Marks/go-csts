package main

import (
	"fmt"
	"log"

	"github.com/2marks/csts/cmd/api"
	"github.com/2marks/csts/config"
	"github.com/2marks/csts/database"
)

func main() {
	newDatabase := database.NewDatabase()
	db, err := newDatabase.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	apiServer := api.NewApiServer(db, fmt.Sprintf(":%s", config.Envs.Port))
	if err = apiServer.Run(); err != nil {
		log.Fatal(err)
	}
}
