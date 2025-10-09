package hmlandreg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var db *sql.DB

func Init() {
	fmt.Println(os.Getenv("DB_ADDR"))
	dsn := mysql.NewConfig()
	dsn.Net = "tcp"
	dsn.User = os.Getenv("DB_USER")
	dsn.Passwd = os.Getenv("DB_PASS")
	dsn.Addr = os.Getenv("DB_ADDR")
	dsn.DBName = os.Getenv("DB_NAME")

	var err error
	db, err = sql.Open("mysql", dsn.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("db connected")
}

func TestCall() {
	var prices []int
	q := sq.Select("price").From("house_sales")

	rows, err := q.RunWith(db).Query()
	if err != nil {
		fmt.Println("query error")
	}

	for rows.Next() {
		var sale int
		if err := rows.Scan(&sale); err != nil {
			fmt.Println("row scan error")
		}

		prices = append(prices, sale)

	}

	fmt.Println(prices)
}
