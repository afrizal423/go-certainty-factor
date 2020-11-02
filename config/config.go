package config

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

// kita buat koneksi dgn db posgres
func CreateConnection() (*sqlx.DB, error) {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db, err := sqlx.Connect("mysql", os.Getenv("MYSQL_CONN"))

	if err != nil {
		panic(err)
	}

	// fmt.Println("Sukses Konek ke Db!")
	// return the connection
	return db, nil
}
