package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create creates a new record in the database
func Create(c *gin.Context) {
	company := Company{}
	company.SetProfileID()
	// company.FundingDetails = &FundingDetails{}
	dbc.Create(company)
	c.JSON(http.StatusOK, company)
}
