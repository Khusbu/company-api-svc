package main

import "github.com/dchest/uniuri"

//GenerateToken generates a random token of 16 characters
func GenerateToken() string {
	return uniuri.New()
}
