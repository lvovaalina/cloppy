package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Repo struct {
	db *sql.DB
}

func NewRepo() Repo {
	return Repo{}
}

func (r *Repo) ConnectDB() error {
	log.Println("LOAD DB, DB PATH: ", DB_PATH)
	db, err := sql.Open("sqlite3", DB_PATH)
	if err != nil {
		log.Println("Failed to load DB: ", err)

		return err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	r.db = db
	return nil
}

func (r *Repo) Save(val []byte) {
	now := time.Now().UTC()
	date := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
	datetimeNow := fmt.Sprintf("%s %d:%d:%d", date, now.Hour(), now.Minute(), now.Second())
	data := escape(string(val))
	insertQuery := fmt.Sprintf(
		"insert into \"history\"(\"value\", \"insert_datetime\") values('%s', '%s')",
		data, datetimeNow)

	log.Println(insertQuery)
	_, e := r.db.Exec(insertQuery)
	checkError(e)
}

func (r *Repo) GetValues(numberOfValues int) []string {
	log.Println("Retrieving last 10 values")
	result := make([]string, 0, 10)
	log.Println("DB ", r.db)
	rows, err := r.db.Query("SELECT value FROM history ORDER BY id DESC;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var val string
		if err := rows.Scan(&val); err != nil {
			log.Fatal(err)
		}

		result = append(result, val)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func (r *Repo) CloseConnection() {
	r.db.Close()
}

func escape(value string) string {
	return strings.ReplaceAll(value, "'", "''")
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
