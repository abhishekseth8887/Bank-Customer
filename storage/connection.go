package storage

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {

	log.Println(logtag + " [dbConn] Started")

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "customers"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(logtag + " [dbConn] Finished")

	return db
}

func dbConnRDS() (db *sql.DB) {

	log.Println(logtag + " [dbConnRDS] Started")

	dbDriver := "mysql"
	dbUser := "admin"
	dbPass := "#Admin123456"
	dbName := "(rds-1.chpbm7kawk5u.ap-southeast-1.rds.amazonaws.com)/customers"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp"+dbName)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(logtag + " [dbConnRDS] Finished")
	return db
}


