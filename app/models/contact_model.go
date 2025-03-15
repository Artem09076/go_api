package models

import "time"

type Contact struct {
	ID         int
	Username   string
	GivenName  string
	FamilyName string
	Phone      []struct {
		TypeID      int
		CountryCode int
		Operator    int
		Number      int
	}
	Email     []string
	Birthdate time.Time
}
