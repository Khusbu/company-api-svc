package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	internalError  = map[string]string{"error": "something didn't work on our side"}
	recordNotFound = map[string]string{"error": "uh oh! company not found"}
	badRequest     = map[string]string{"error": "please send data in proper JSON format"}
)

// Create creates a new record in the database
func Create(c *gin.Context) {
	company, err := ConvertPostDatatoModel(c)
	if err != nil {
		log.Printf("Error in retrieving post data: %q", err)
		c.JSON(http.StatusBadRequest, badRequest)
		return
	}
	if err := company.Validate(); err != nil {
		log.Printf("Invalid field format: %q", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := CreateRecord(company); err != nil {
		log.Printf("Error creating record: %q", err)
		c.JSON(http.StatusInternalServerError, internalError)
		return
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
			c.JSON(http.StatusNotFound, recordNotFound)
			return
		}
		log.Printf("Error fetching record: %q", err)
		c.JSON(http.StatusInternalServerError, internalError)
		return
	}
	c.JSON(http.StatusOK, company)
}
