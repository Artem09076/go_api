package models

import (
	"github.com/google/uuid"
)

type Phone struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	CountyCode int        `db:"countycode"`
	Operator   int        `db:"operator"`
	Number     int        `db:"number"`
	ContactID  *uuid.UUID `gorm:"type:uuid" json:"-"`
}

type Contact struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Username     string     `db:"username" json:"username"`
	GivenName    string     `db:"given_name" json:"given_name"`
	Email        string     `db:"email" json:"email"`
	Phones       []Phone    `gorm:"foreignKey:ContactID" json:"phones"`
	BirthdayDate string     `db:"birthday_date" json:"birthday_date"`
	GroupID      *uuid.UUID `db:"group_id" json:"-"`
}
