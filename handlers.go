package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Create creates a new record in the database
func Create(c *gin.Context) {
	company := Company{}
	company.SetProfileID()
	company.FundingDetails = &FundingDetails{Amount: 123}
	if err := CreateRecord(company); err != nil {
		log.Printf("Error creating record: %q", err)
		c.JSON(http.StatusInternalServerError, nil) //TODO pass some data
	}
	c.JSON(http.StatusOK, company)
}

// Fetch fetches company details from the database with the requested profile_id
func Fetch(c *gin.Context) {
	profileID := c.Param("profile_id")
	company, err := FetchRecord(profileID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Record not found: %s", profileID)
			c.JSON(http.StatusNotFound, nil)
			return
		}
		log.Printf("Error fetching record: %q", err)
		c.JSON(http.StatusInternalServerError, nil) //TODO pass some data
		return
	}
	c.JSON(http.StatusOK, company)
}
