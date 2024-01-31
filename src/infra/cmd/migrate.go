package main

import (
	"dating-services/src/infra"
	"dating-services/src/infra/database"
	"dating-services/src/infra/persistence/postgres"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := infra.Make()
	db := postgres.New(config.SqlDb)

	// migration
	err := database.Migrate(db.DB)
	if err != nil {
		panic(err)
	}

	log.Println("Success migration")
}
