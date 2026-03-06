package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	_ = godotenv.Load()

	addr := os.Getenv("DB_ADDR")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	mode := os.Getenv("SSL_MODE")

	conf := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", addr, user, pass, name, port, mode)
	db, err := sql.Open("postgres", conf)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
