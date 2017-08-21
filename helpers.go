package main

import (
	"errors"

	"github.com/dchest/uniuri"
	"github.com/gin-gonic/gin"
)

const (
	urlEncodedHeader = "application/x-www-form-urlencoded"
	jsonHeader       = "application/json"
)

// ConvertPostDatatoModel retrieves post params
func ConvertPostDatatoModel(c *gin.Context) (company Company, err error) {
	switch c.ContentType() {
	case jsonHeader:
		company.SetProfileID()
		err = c.Bind(&company)
	default:
		err = errors.New("Only JSON format supported")
	}
	return
}

//GenerateToken generates a random token of 16 characters
func GenerateToken() string {
	return uniuri.New()
}

// IsIn check if string str is a member of the set of strings params
func IsIn(str string, params ...string) bool {
	for _, param := range params {
		if str == param {
			return true
		}
	}

	return false
}
