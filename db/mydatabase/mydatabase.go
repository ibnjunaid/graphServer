package mydatabase

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	//idk why its there
	_ "github.com/lib/pq"

	//idk why its there
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//Db is unjust
var Db *sql.DB

//InitDB initialises a connection to the database
func InitDB() {
	connStr := "user=osama dbname=hackernews password=ibnjunaid "
	// Use root:dbpass@tcp(172.17.0.2)/hackernews, if you're using Windows.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}
	Db = db

}

//Migrate creates table if it doesnt exist
func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := postgres.WithInstance(Db, &postgres.Config{})

	m, e := migrate.NewWithDatabaseInstance(
		"file:///home/osama/go/src/graphServer/db/mydatabase",
		"postgres",
		driver,
	)
	if e != nil {
		log.Fatal(e)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("Eror")
		log.Fatal(err)
	}

}
