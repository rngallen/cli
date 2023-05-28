package data

import (
	"fmt"
	"log"
	"os"

	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func ConnectDb() {
	var err error

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable timezone=Africa/Nairobi",
		"LOCALHOST", 5432, "MY DB USER", "MY DB USER PWD", "MY DB")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Minute,
			LogLevel:      logger.Error,
			Colorful:      true,
		},
	)

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			// 	TablePrefix:   "",
			// 	SingularTable: true, 	     				   // use singular table name, table for `User` would be `user` with this option enabled
			// 	NoLowerCase:   false,          	              // skip the snake_casing of names LastName will be LastName if is false otherwise last_name
			NameReplacer: strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	sqlDB, err := Db.DB()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Maximum Iddle Connections
	sqlDB.SetMaxIdleConns(10)
	// Maximum Open Connections
	sqlDB.SetMaxOpenConns(100)
	// Maximum time a connection can be re-used
	sqlDB.SetConnMaxLifetime(time.Hour)

}
