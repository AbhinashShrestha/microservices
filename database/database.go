package database

import {
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/AbhinashShrestha/rest_api/models"
} 

type Dbinstance struct{
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb(){
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Kathmandu",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME")
	)

	db , error := gorm.Open(postgres.Open(dsn),&gorm.Config{
		logger:logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to the database. \n",err)
		os.Exit(2)
	}

	log.Println("Sucessfully connected to the postgres database")
	db.logger = logger.Default.LogMode(logger.Info)

	log.Println("running migration")
	db.AutoMigrate(&.models.Fact{})

	DB = Dbinstance {
		Db: db,
	}
}