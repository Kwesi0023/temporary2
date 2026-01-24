package initializers

import (
	"log"
	"os"

	//	gorm.io/driver/mongodb // gorm does not support mongodb, i think for now
	// they only support MySQL, PostgreSQL, SQLite, SQL Server, and others.
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mongodb.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to Connect to Database")
	}

}
