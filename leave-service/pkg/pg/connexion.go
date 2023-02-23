package pg

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// counts holds the number of connexion attempt to the db
var counts int64

// PostgreSQL Credentials
const (
	host     = "backend"
	port     = 5432
	user     = "postgres"
	password = "rootroot"
	dbname   = "authentication"
)

// openDB opens a postgres database for driver-specific data source name
func openDB(dsn string) (*sql.DB, error) {
	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ConnectToDB try to connect endlessly to postgres database
// it sleeps for 2 seconds when 10 attempts failed
func ConnectToDB() *sql.DB {
	// dsn := os.Getenv("DSN")
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}
