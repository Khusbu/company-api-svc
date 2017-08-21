package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

// Company defines structure of company details
type Company struct {
	Name           string            `json:"name" gorm:"not null" valid:"length(0|20)"`
	Description    string            `json:"description" valid:"length(0|200)"`
	Logo           string            `json:"logo" valid:"url,length(0|100)"`
	FundingDetails []*FundingDetails `json:"funding_details" gorm:"ForeignKey:ProfileID;AssociationForeignKey:ProfileID"`
	Markets        string            `json:"markets" valid:"length(0|100)"`
	FoundedOn      string            `json:"founded_on" valid:"matches(^(?:0?[1-9]|[1-2][0-9]|3[01])/(?:0?[1-9]|1[0-2])/[0-9]{4}$)"`
	Website        string            `json:"website" valid:"url,length(0|100)"`
	SocialInformation
	ProfileID string `json:"profile_id" gorm:"primary_key;index"`
}

// FundingDetails defines structure of funding details of a company
type FundingDetails struct {
	ID        uint   `gorm:"primary_key"`
	Amount    string `json:"amount" gorm:"not null" valid:"numeric,length(0|10)"`
	Date      string `json:"date" valid:"matches(^(?:0?[1-9]|[1-2][0-9]|3[01])/(?:0?[1-9]|1[0-2])/[0-9]{4}$)"`
	Stages    string `json:"stages"`
	Investors string `json:"investors" valid:"length(0|100)"`
	ProfileID string `json:"profile_id"`
}

// SocialInformation defines structure of social information of a company
type SocialInformation struct {
	LinkedIn    string `json:"linked_in" valid:"url,length(0|100)"`
	Twitter     string `json:"twitter" valid:"url,length(0|100)"`
	Email       string `json:"email" valid:"email,length(0|100)"`
	PhoneNumber string `json:"phone_number" valid:"numeric,length(10|10)"`
}

// SetProfileID sets new ProfileID of a company
func (company *Company) SetProfileID() {
	company.ProfileID = "cmp_" + GenerateToken()
}

// Validate validates the fields of a company model
func (company *Company) Validate() error {
	if _, err := govalidator.ValidateStruct(company); err != nil {
		return err
	}
	if _, err := govalidator.ValidateStruct(company.SocialInformation); err != nil {
		return err
	}
	for _, fd := range company.FundingDetails {
		if err := fd.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Validate validates the fields of a FundingDetails model
func (fundingDetails *FundingDetails) Validate() error {
	if _, err := govalidator.ValidateStruct(fundingDetails); err != nil {
		return err
	}
	stages := []string{"Series A", "Series B", "Series C", "Series D", "Series E", "Series F"}
	if !IsIn(fundingDetails.Stages, stages...) {
		return fmt.Errorf("Stages: Got %s", fundingDetails.Stages)
	}
	return nil
}

// Validate validates the fields of a SocialInformation model
func (socialInformation *SocialInformation) Validate() error {
	if _, err := govalidator.ValidateStruct(socialInformation); err != nil {
		return err
	}
	return nil
}
