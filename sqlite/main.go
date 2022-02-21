package main

import(
	"fmt"
	"runner/db"
	"log"
	"time"
	"database/sql"
)

type Timestamp time.Time

func main(){
	filepath := "../data/db.db"

	err := db.CreateDBFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	var database *sql.DB

	database, err = db.InitialDB(filepath)
	if err != nil {
		db.CloseDB(database)
		log.Fatal(err)
	}

	now_time := time.Now()

	database, err = db.InsertData(
		database, 
		"FX6e-MC 1호기",
		"192.168.0.1",
		4547,
		"광교",
		now_time,
		true,
		false,
		true,
		false,
		true,
	)
	if err != nil {
		db.CloseDB(database)
		log.Fatal(err)
	}

	db.PrintAllData(database)

	err = db.CloseDB(database)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Run Success.")
}