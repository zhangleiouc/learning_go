package bootstrap

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLDatabase(env *Env) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DBUser,
		env.DBPass,
		env.DBHost,
		env.DBPort,
		env.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database: ", err)
	}

	log.Println("Connected to MySQL database successfully")
	return db
}

func CloseMySQLConnection(db *sql.DB) {
	if db == nil {
		return
	}

	err := db.Close()
	if err != nil {
		log.Fatal("Failed to close database connection: ", err)
	}

	log.Println("Connection to MySQL closed.")
}
