package main

import (
	"log"

	"github.com/bimmagung/flight-ticket-api/pkg/db/postgresdb"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	connect :=  postgresdb.NewPostgresDB()

	if connect == nil {
		log.Panic("can't connect to postgres database")
	}

	defer func ()  {
		err := connect.Close()
		if err != nil {
			log.Panic(err)
		}
	}()
}