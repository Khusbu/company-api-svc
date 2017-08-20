package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create creates a new record in the database
func Create(c *gin.Context) {
	company := Company{}
	company.SetProfileID()
	company.FundingDetails = &FundingDetails{}
	dbc.Create(company)
	c.JSON(http.StatusOK, company)
}

// Fetch fetches company details from the database with the requested profile_id
func Fetch(c *gin.Context) {
	profileID := c.Param("profile_id")
	var company Company
	dbc.Preload("FundingDetails").Where("profile_id = ?", profileID).First(&company)
	c.JSON(http.StatusOK, company)
}
