package storage

import (
	"database/sql"
	"log"
	"myapps/bank/entity"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "customers"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Insert(customer entity.Customer) {

	db := dbConn()

	insForm, err := db.Prepare(
		"INSERT INTO Customers(" +
			"first_name, middle_name,last_name, dob, mobile_number, gender," +
			"customer_number, country_of_birth, country_of_residence, customer_segment) " +
			"VALUES(?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	insForm.Exec(customer.FirstName, customer.MiddleName, customer.LastName,
		customer.Dob, customer.MobileNumber, customer.Gender,
		customer.CustomerNumber, customer.CountryOfBirth, customer.CountryOfResidence,
		customer.CustomerSegment)

	defer db.Close()
	log.Println("INSERT: ")
}
