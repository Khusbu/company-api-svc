package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var (
	dbc *gorm.DB
)

// CreateTables creates tables for each model //TODO add error handling
func CreateTables() {
	dbc.CreateTable(&Company{}, &FundingDetails{})
	dbc.Model(&FundingDetails{}).AddForeignKey("profile_id", "companies(profile_id)", "RESTRICT", "RESTRICT")
	dbc.AutoMigrate(&Company{}, &FundingDetails{})
}

// CreateRecord creates a company record in the database
func CreateRecord(company Company) error {
	return dbc.Create(&company).Error
}

// FetchRecord fetches a company record from the database having profile_id `id`
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
	dbc.LogMode(true)
	// CreateTables() //TODO uncomment
}
