package main

// CreateTables creates tables for each model
func CreateTables() {
	dbc.CreateTable(&Company{}, &FundingDetails{})
	dbc.Model(&FundingDetails{}).AddForeignKey("profile_id", "companies(profile_id)", "RESTRICT", "RESTRICT")
	dbc.AutoMigrate(&Company{}, &FundingDetails{})
}
