package utils

import (
	"database/sql"
	"fmt"

	"github.com/DATA-DOG/go-txdb"
	"github.com/bgw7/products-api/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func RegisterDatabase() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("error loading .env file")
	}

	config.Init()
	cfg := config.GetConfig()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
	)
	txdb.Register("txdb", "mysql", dsn)
}

func InitDatabase() error {
	conn, err := sql.Open("txdb", "identifier")

	db = conn

	return err
}

func GetDB() *sql.DB {
	return db
}
