package db

import (
	"database/sql"
	_ "errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

type MsgData struct {
	Location  string `json:"loc"`
	Timestamp string `json:"time"`
	IsRun     bool   `json:"IsRun"`
}

/*********************************
	Create Physical db file
***********************************/
func CreateDBFile(filepath string) error {
	os.Remove(filepath)
	log.Println("[CRIT] Removed " + filepath + " file.")

	file, err := os.Create(filepath) // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	log.Println("[CRIT] Create " + filepath + " file.")

	return err
}

/*********************************
	Create DB table
***********************************/
func InitialDB(filepath string) (*sql.DB, error) {
	sqliteDatabase, _ := sql.Open("sqlite3", filepath)

	createTesterStatusTableSQL := `CREATE TABLE tester_status (
		"idx" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"equip_name" TEXT,
		"equip_ip" TEXT,
		"equip_port" INT,
		"location" TEXT,
		"report_time" TIMESTAMP,
		"stat_pc" BOOLEAN,
		"stat_prog_tester" BOOLEAN,
		"stat_prog_chamber" BOOLEAN,
		"stat_prog_cloud" BOOLEAN,
		"stat_prog_test_manager" BOOLEAN
	);`

	statement, err := sqliteDatabase.Prepare(createTesterStatusTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("[MSG ] teser_status table created")

	return sqliteDatabase, err
}

func InsertData(
	db *sql.DB,
	equip_name string,
	equip_ip string,
	equip_port int,
	location string,
	report_time time.Time,
	stat_pc bool,
	stat_prog_tester bool,
	stat_prog_chamber bool,
	stat_prog_cloud bool,
	stat_prog_test_manager bool,
) (*sql.DB, error) {
	log.Println(equip_name, equip_ip, equip_port, location, report_time, stat_pc, stat_prog_tester, stat_prog_chamber, stat_prog_cloud, stat_prog_test_manager)
	insertTesterStatusSQL := `INSERT INTO tester_status(
			equip_name, 
			equip_ip, 
			equip_port, 
			location, 
			report_time,
			stat_pc, 
			stat_prog_tester, 
			stat_prog_chamber, 
			stat_prog_cloud, 
			stat_prog_test_manager
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertTesterStatusSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(
		equip_name,
		equip_ip,
		equip_port,
		location,
		report_time,
		stat_pc,
		stat_prog_tester,
		stat_prog_chamber,
		stat_prog_cloud,
		stat_prog_test_manager,
	)
	if err != nil {
		log.Fatalln(err.Error())
	}

	
	return db, err
}

func PrintAllData(db *sql.DB){
	row, err := db.Query("SELECT * FROM tester_status ORDER BY equip_name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var idx int
		var equip_name string
		var equip_ip string
		var equip_port int
		var location string
		var report_time time.Time
		var stat_pc bool
		var stat_prog_tester bool
		var stat_prog_chamber bool
		var stat_prog_cloud bool
		var stat_prog_test_manager bool
		row.Scan(
			&idx,
			&equip_name,
			&equip_ip,
			&equip_port,
			&location,
			&report_time,
			&stat_pc,
			&stat_prog_tester,
			&stat_prog_chamber,
			&stat_prog_cloud,
			&stat_prog_test_manager,	
		)
		log.Println("LSIT =========================")
		log.Println(
			equip_name,
			equip_ip,
			equip_port,
			location,
			report_time.Format("2006.01.02 15:04:05"),
			stat_pc,
			stat_prog_tester,
			stat_prog_chamber,
			stat_prog_cloud,
			stat_prog_test_manager,
		)
		log.Println("==============================")
	}

}

func CloseDB(sqliteDatabase *sql.DB) error{
	err := sqliteDatabase.Close() // Defer Closing the database
	log.Println("Database Closed.")
	return err
}