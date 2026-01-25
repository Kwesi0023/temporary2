package initializers

import (
	"log"
	//	gorm.io/driver/mongodb // gorm does not support mongodb, i think for now
	// they only support MySQL, PostgreSQL, SQLite, SQL Server, and others.
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=localhost user=postgres password=00000 dbname= port=5432 sslmode=disable"
	// dbname= "YOUR_DB_NAME"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to Connect to Database: ", err)
	}

}
