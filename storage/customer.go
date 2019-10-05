package storage

import (
	"log"
	"myapps/Bank-Customer/entity"
)

func Get() (*[]entity.Customer, error) {

	log.Println(logtag + " [Get] Started")

	db := dbConn()
	//db := dbConnRDS()
	defer db.Close()

	customers := []entity.Customer{}

	stmtOut, err := db.Prepare(
		"SELECT * FROM customers order by created_at desc limit 100")

	if err != nil {
		log.Println(err.Error())
		return &customers, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Println(err.Error())
		return &customers, err
	}

	for rows.Next() {
		var customer entity.Customer
		var resAdd, offAdd entity.Address
		var id int64
		var created_at, updated_at []uint8

		resAdd.Type = "Residence"
		offAdd.Type = "Office"
		err = rows.Scan(&id, &customer.FirstName, &customer.MiddleName, &customer.LastName,
			&customer.Dob, &customer.MobileNumber, &customer.Gender, &customer.CustomerNumber,
			&customer.CountryOfBirth, &customer.CountryOfResidence, &customer.CustomerSegment,
			&resAdd.AddrLine1, &resAdd.AddrLine2, &resAdd.AddrLine3, &resAdd.AddrLine4,
			&resAdd.ZipCode, &resAdd.City, &resAdd.State, &resAdd.CountryCode,
			&offAdd.AddrLine1, &offAdd.AddrLine2, &offAdd.AddrLine3, &offAdd.AddrLine4,
			&offAdd.ZipCode, &offAdd.City, &offAdd.State, &offAdd.CountryCode,
			&created_at, &updated_at,
		)
		customer.Addresses = []entity.Address{resAdd, offAdd}

		if err != nil {
			log.Println(err.Error())
			continue
		}
		customers = append(customers,customer)
	}

	if err != nil {
		log.Println(err.Error())
		return &customers, err
	}

	log.Println(logtag + " [Get] Finished")

	return &customers, err
}

func Insert(customer entity.Customer) error{

	log.Println(logtag + " [Insert] Started")

	db := dbConn()
	//db := dbConnRDS()
	defer db.Close()

	query, err := db.Prepare(
		"INSERT INTO customers(" +
			"first_name, middle_name, last_name, dob, " +
			"mobile_number, gender, customer_number, country_of_birth, " +
			"country_of_residence, customer_segment," +

			"res_addr_line_1, res_addr_line_2, res_addr_line_3, res_addr_line_4," +
			"res_zip_code, res_city, res_state, res_country_code," +
			"off_addr_line_1, off_addr_line_2, off_addr_line_3, off_addr_line_4," +
			"off_zip_code, off_city, off_state, off_country_code)" +

			"VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = query.Exec(customer.FirstName, customer.MiddleName, customer.LastName,
		customer.Dob, customer.MobileNumber, customer.Gender,
		customer.CustomerNumber, customer.CountryOfBirth, customer.CountryOfResidence,
		customer.CustomerSegment,

		customer.Addresses[0].AddrLine1, customer.Addresses[0].AddrLine2,
		customer.Addresses[0].AddrLine3, customer.Addresses[0].AddrLine4,
		customer.Addresses[0].ZipCode, customer.Addresses[0].City,
		customer.Addresses[0].State, customer.Addresses[0].CountryCode,

		customer.Addresses[1].AddrLine1, customer.Addresses[1].AddrLine2,
		customer.Addresses[1].AddrLine3, customer.Addresses[1].AddrLine4,
		customer.Addresses[1].ZipCode, customer.Addresses[1].City,
		customer.Addresses[1].State, customer.Addresses[1].CountryCode)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(logtag + " [Insert] Finished")

	return err
}

func Update(customer entity.Customer) error {

	log.Println(logtag + " [Update] Started")

	db := dbConn()
	//db := dbConnRDS()
	defer db.Close()

	query, err := db.Prepare(
		"UPDATE customers SET dob=?, mobile_number=?, gender=?, customer_number=?, " +
			"country_of_birth=?, country_of_residence=?, customer_segment=?," +
			"res_addr_line_1=?, res_addr_line_2=?, res_addr_line_3=?, res_addr_line_4=?," +
	        "res_zip_code=?, res_city=?, res_state=?, res_country_code=?," +
			"off_addr_line_1=?, off_addr_line_2=?, off_addr_line_3=?, off_addr_line_4=?," +
			"off_zip_code=?, off_city=?, off_state=?, off_country_code=?" +
			" WHERE first_name=? and middle_name=? and last_name=?")

	if err != nil {
		return err
		log.Println(err.Error())
	}

	_, err = query.Exec(customer.Dob, customer.MobileNumber, customer.Gender,
		customer.CustomerNumber, customer.CountryOfBirth, customer.CountryOfResidence,
		customer.CustomerSegment,
		customer.Addresses[0].AddrLine1, customer.Addresses[0].AddrLine2,
		customer.Addresses[0].AddrLine3, customer.Addresses[0].AddrLine4,
		customer.Addresses[0].ZipCode, customer.Addresses[0].City,
		customer.Addresses[0].State, customer.Addresses[0].CountryCode,

		customer.Addresses[1].AddrLine1, customer.Addresses[1].AddrLine2,
		customer.Addresses[1].AddrLine3, customer.Addresses[1].AddrLine4,
		customer.Addresses[1].ZipCode, customer.Addresses[1].City,
		customer.Addresses[1].State, customer.Addresses[1].CountryCode,

		customer.FirstName, customer.MiddleName, customer.LastName)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(logtag + " [Insert] Finished")

	return err
}

func Delete(customer entity.Customer) error {

	log.Println(logtag + " [Delete] Started")

	db := dbConn()
	//db := dbConnRDS()
	defer db.Close()

	query, err := db.Prepare(
		"DELETE from customers " +
			" WHERE first_name=? and middle_name=? and last_name=?")

	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = query.Exec(customer.FirstName, customer.MiddleName, customer.LastName)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(logtag + " [Delete] Finished")

	return err
}
