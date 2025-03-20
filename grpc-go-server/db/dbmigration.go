package dbmigration

import (
	"database/sql"
	"log"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(conn *sql.DB) {
	log.Println("Database migration start")

	driver, err := postgres.WithInstance(conn, &postgres.Config{})

	if err != nil{
		log.Fatalln("Failed to get Postgres Instance: ", err)
	}
	
	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "grpc", driver)

	if err != nil {
		log.Println("Database migration failed :", err)
	}

	if err := m.Down(); err != nil {
		log.Println("Database migration (down) failed :", err)
	}

	if err := m.Up(); err != nil {
		log.Println("Database migration (up) failed :", err)
	}

	log.Println("Database migration end")
}
