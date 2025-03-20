package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4/stdlib"
	dbmigration "github.com/tanalam2411/grpc-demo/db"

	db "github.com/tanalam2411/grpc-demo/internal/adapter/database"
	grpc "github.com/tanalam2411/grpc-demo/internal/adapter/grpc"
	app "github.com/tanalam2411/grpc-demo/internal/application"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	sqlDB, err := sql.Open("pgx", "postgres://postgres:admin@localhost:5432/grpc?sslmode=disable")

	if err != nil{
		log.Fatalln("Can't connect to database: ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalln("Connection failed: ", err)
	}

	dbmigration.Migrate(sqlDB)

	databaseAdapter, err := db.NewDatabaseAdapter(sqlDB)

	if err != nil {
		log.Fatalln("Can't create database adapter: ", err)
	}

	// runDummyOrm(databaseAdapter)

	hs := &app.HelloService{}
	bs := app.NewBankService(databaseAdapter)

	grpcAdapter := grpc.NewGrpcAdapter(hs, bs, 9090)

	grpcAdapter.Run()
}

func runDummyOrm(da *db.DatabaseAdapter) {
	now := time.Now()

	uuid, _ := da.Save(
		&db.DummyOrm{
			UserId: uuid.New(),
			UserName: "Tim " + time.Now().Format("15:04:05"),
			CreatedAt: now,
			UpdatedAt: now,
		},
	)

	res, _ := da.GetByUuid(&uuid)

	log.Println("res: ", res)
}