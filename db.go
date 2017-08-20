package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var (
	dbc *gorm.DB
)

// CreateTables creates tables for each model
func CreateTables() {
	dbc.CreateTable(&Company{}, &FundingDetails{})
	dbc.Model(&FundingDetails{}).AddForeignKey("profile_id", "companies(profile_id)", "RESTRICT", "RESTRICT")
	dbc.AutoMigrate(&Company{}, &FundingDetails{})
}

// CreateRecord creates a record in the database
func CreateRecord(model interface{}) error {
	return dbc.Create(model).Error
}

// FetchRecord fetches a record from the database having profile_id `id`
func FetchRecord(id string) (Company, error) {
	var company Company
	if err := dbc.Preload("FundingDetails").Where("profile_id = ?", id).First(&company).Error; err != nil {
		return company, err
	}
	return company, nil
}

// SetUpDB establishes a connection to the database and create tables
func SetUpDB() {
	var err error
	dbc, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	CreateTables()
}
