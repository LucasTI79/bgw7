package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bgw7/products-api/config"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// dsn = "user:password@(tcp:host)/dbname"
func Initialize() error {
	var err error
	cfg := config.GetConfig()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
	)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	log.Println("database connect succefully")

	return nil
}

func GetDB() *sql.DB {
	return db
}
