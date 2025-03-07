package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME_AUTH")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	var dsn string

	if password == "" {
		// Jika password kosong, tidak sertakan parameter password
		dsn = fmt.Sprintf(
			"host=%s user=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Makassar",
			host, username, database, port, "disable",
		)
	} else {
		// Jika password ada, sertakan parameter password
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Makassar",
			host, username, password, database, port, "disable",
		)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Second * time.Duration(300))

	return db
}
