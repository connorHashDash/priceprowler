package hmlandreg

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

var db *gorm.DB

func Init() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_ADDR"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("can't connect to DB")
	}

	var version string
	if err := db.Raw("SELECT VERSION()").Scan(&version).Error; err != nil {
		log.Fatal("test query failed: ", err)
	}

	fmt.Println(version)

}
