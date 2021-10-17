package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "cloppy"
)

var db *sql.DB

func connectDB() {
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error

	db, err = sql.Open("postgres", psqlconn)
	checkError(err)
}

func add(value []byte) {
	log.Println(time.Now().UTC().UnixNano())
	now := time.Now().Local()
	date := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
	datetimeNow := fmt.Sprintf("%s %d:%d:%d", date, now.Hour(), now.Minute(), now.Second())
	data := escape(string(value))
	insertQuery := fmt.Sprintf(
		"insert into \"clipboardhistory\"(\"data\", \"insert_datetime\") values('%s', '%s')",
		data, datetimeNow)

	log.Println(insertQuery)
	_, e := db.Exec(insertQuery)
	checkError(e)
}

func getValues(numberOfValues int) []string {
	log.Println("Retrieving last 10 values")
	result := make([]string, 0)
	rows, err := db.Query("SELECT data FROM clipboardhistory ORDER BY id DESC;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			log.Fatal(err)
		}
		fmt.Println(title)
		result = append(result, title)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func escape(value string) string {
	return strings.ReplaceAll(value, "'", "''")
}

func close() {
	db.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
