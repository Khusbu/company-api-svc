package main

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

// Company defines structure of company details
type Company struct {
	Name           string            `json:"name" gorm:"not null"`
	Description    string            `json:"description" valid:"length(0|200)"`
	Logo           string            `json:"logo" valid:"url,length(0|100)"`
	FundingDetails []*FundingDetails `json:"funding_details" gorm:"ForeignKey:ProfileID;AssociationForeignKey:ProfileID"`
	Markets        string            `json:"markets" valid:"length(0|100)"`
	FoundedOn      string            `json:"founded_on" valid:"numeric,length(0|10)"` //TODO assume int64
	Website        string            `json:"website" valid:"url,length(0|100)"`
	SocialInformation
	ProfileID string `json:"profile_id" gorm:"primary_key;index"`
}

// FundingDetails defines structure of funding details of a company
type FundingDetails struct {
	ID        uint   `gorm:"primary_key"`
	Amount    string `json:"amount" valid:"numeric,length(0|10)"`                                 //TODO assume paise
	Date      string `json:"date" valid:"numeric,length(0|10)"`                                   //TODO assume unix time
	Stages    string `json:"stages" valid:"in('SeriesA|SeriesB|SeriesC|SeriesD|SeriesE|SeriesF)"` //TODO assume single value
	Investors string `json:"investors" valid:"length(0|100)"`                                     //TODO assume single value
	ProfileID string `json:"profile_id"`
}

// SocialInformation defines structure of social information of a company
type SocialInformation struct {
	LinkedIn    string `json:"linked_in" valid:"url,length(0|100)"`
	Twitter     string `json:"twitter" valid:"url,length(0|100)"`
	Email       string `json:"email" valid:"email,length(0|100)"`
	PhoneNumber string `json:"phone_number" valid:"numeric,length(10|10)"` //TODO assume 10 digit mobile number
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
		return errors.New("Invalid value for stages")
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
