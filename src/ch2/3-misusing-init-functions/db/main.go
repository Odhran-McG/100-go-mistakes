package main

import (
	"database/sql"
	"log"
	"os"
)

// Global variable
//		any package function can alter this value
//		unit tests that test around this are no longer isolated due to prior point
var db *sql.DB

// Opening connection pool can fail
//	this means the package init function can close the application - not the caller
//		maybe okay for some logic but better to make this decision explicitly elsewhere
//	caller doesnt get a chance to retry or implement fallback logic as init doesn't return an error
//	init is called before all tests, couples db setup to all tests, even unrelated unit tests
func init() {
	dataSourceName := os.Getenv("MY_SQL_DATASOURCE_NAME")
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	err = d.Ping()
	if err != nil {
		log.Panic(err)
	}
	db = d
}

// Alternative to using init in this case
//		error handling responibilty of caler
//		easily test with an integration test
//		connection pool, db pointer, is encapsulated within the function
func createClient(dataSourceName string)  (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil

}

func main(){
	createClient("var")
}