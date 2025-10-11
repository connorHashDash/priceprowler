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

func GetPriceChange_AllTypes() ([]PriceTrendData, error) {
	var TrendData []PriceTrendData
	q := sq.Select(
		"DATE_FORMAT(transfer_date, '%Y-%m') AS month",
		"property_type",
		"ROUND(AVG(price)) AS avg_price").
		From("house_sales").
		Where(sq.Eq{"record_status": "h"}).
		GroupBy("month", "property_type").
		OrderBy("property_type")

	rows, err := q.RunWith(db).Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var r PriceTrendData
		if err := rows.Scan(&r.Month, &r.PropertyType, &r.AvgPrice); err != nil {
			return nil, err
		}
		TrendData = append(TrendData, r)
	}

	return TrendData, nil
}
