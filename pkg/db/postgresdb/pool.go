package postgresdb

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var counts int64

func openDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}

func NewPostgresDB() *sqlx.DB {

	
	var (
		user = os.Getenv("DBUSER")
		password = os.Getenv("DBPASSWORD")
		dbname = os.Getenv("DBNAME")
		host = os.Getenv("DBHOST")
		port = os.Getenv("DBPORT")
		sslmode = os.Getenv("DBSSLMODE")
	)

	dsn :=  fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", 
	host, port, user, password, dbname, sslmode)

	for {
		conn, err := openDB(dsn)

		if err != nil {
			log.Println("postgres not yet ready...")
			counts++
		}else {
			log.Println("connected to postgress!")
			return conn
		}

		if  counts > 10 {
			log.Println(err)
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}
