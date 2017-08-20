package main

import "github.com/jinzhu/gorm"

// Company defines structure of company details
type Company struct {
	gorm.Model
	Name           string
	ProfileID      string          `gorm:"primary_key;index"`
	FundingDetails *FundingDetails `gorm:"ForeignKey:ProfileID;AssociationForeignKey:ProfileID"`
}

// FundingDetails defines structure of funding details of a company
type FundingDetails struct {
	gorm.Model
	Amount    int64
	ProfileID string
}

// SetProfileID sets new ProfileID of a company
func (c *Company) SetProfileID() {
	c.ProfileID = "cmp_" + GenerateToken()
}
